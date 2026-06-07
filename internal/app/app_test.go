package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestListValuesIncludesFullRegistry(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{"--list-values"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d: %s", code, stderr.String())
	}

	var catalog ValueCatalog
	if err := json.Unmarshal(stdout.Bytes(), &catalog); err != nil {
		t.Fatalf("unmarshal list-values output: %v", err)
	}

	if got, want := len(catalog.Families), len(FamilyRegistry()); got != want {
		t.Fatalf("family registry length mismatch: got %d want %d", got, want)
	}
	if !containsString(catalog.Config.DevType, DevTypeBackend.String()) {
		t.Fatalf("dev-type registry missing %q", DevTypeBackend)
	}
	if !containsString(catalog.Config.OutputFormat, OutputFormatJSON.String()) {
		t.Fatalf("output-format registry missing %q", OutputFormatJSON)
	}
	if !containsString(catalog.Experimental.Provider, "openai-consumer") {
		t.Fatalf("experimental provider registry missing openai-consumer")
	}
	if !containsString(catalog.Experimental.Provider, "claude-consumer") {
		t.Fatalf("experimental provider registry missing claude-consumer")
	}
	expectedDedicated := map[Family]string{
		FamilyPlatformEngineering:    "dedicated:platform_engineering",
		FamilyObservabilityAIRuntime: "dedicated:observability_ai_runtime",
		FamilyDeliveryPreviewOps:     "dedicated:delivery_preview_ops",
		FamilySupplyChainSecurity:    "dedicated:supply_chain_security",
	}
	for _, spec := range catalog.Families {
		renderer, ok := expectedDedicated[spec.Name]
		if !ok {
			continue
		}
		if spec.Renderer != renderer {
			t.Fatalf("expected %s renderer %q, got %q", spec.Name, renderer, spec.Renderer)
		}
		if !spec.Smoke {
			t.Fatalf("expected %s smoke flag to be true", spec.Name)
		}
	}
}

func TestCodeAnalyzerSmokeEvidence(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{
		"--dev-type", DevTypeBackend.String(),
		"--complexity", ComplexityLow.String(),
		"--seed", "11",
		"--output-format", OutputFormatJSON.String(),
	}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d: %s", code, stderr.String())
	}

	events := decodeEvents(t, stdout.Bytes())
	if len(events) < 3 {
		t.Fatalf("expected start, activity, and end events, got %d", len(events))
	}
	activity := events[1]
	if activity.Context.Activity == nil {
		t.Fatalf("expected activity context")
	}
	if activity.Context.Activity.Family != FamilyCodeAnalyzer.String() {
		t.Fatalf("expected code_analyzer activity, got %s", activity.Context.Activity.Family)
	}
	if !activity.Context.Activity.Smoke {
		t.Fatalf("expected smoke activity")
	}
	if !containsString(activity.Context.Evidence, "dedicated-smoke") {
		t.Fatalf("expected dedicated-smoke evidence, got %v", activity.Context.Evidence)
	}
	if len(activity.Context.Traceability) == 0 || activity.Context.Traceability[0].Source != "rust-stakeholder" {
		t.Fatalf("expected rust-stakeholder traceability row, got %#v", activity.Context.Traceability)
	}
}

func TestClassicSixDedicatedRenderers(t *testing.T) {
	cases := []struct {
		name   string
		family Family
		dev    DevType
	}{
		{name: "data_processing", family: FamilyDataProcessing, dev: DevTypeDataScience},
		{name: "jargon", family: FamilyJargon, dev: DevTypeBackend},
		{name: "metrics", family: FamilyMetrics, dev: DevTypeDevOps},
		{name: "network_activity", family: FamilyNetworkActivity, dev: DevTypeFullStack},
		{name: "system_monitoring", family: FamilySystemMonitoring, dev: DevTypeSystemsProgramming},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := renderFamily(RenderContext{
				Config: Config{DevType: tc.dev},
				Family: tc.family,
				Spec:   FamilySpecFor(tc.family),
			})
			if !strings.Contains(result.Message, tc.family.String()) {
				t.Fatalf("expected dedicated renderer output for %s, got %q", tc.family, result.Message)
			}
			if !containsString(result.Evidence, "dedicated-smoke") {
				t.Fatalf("expected dedicated smoke evidence for %s, got %v", tc.family, result.Evidence)
			}
			if !containsString(result.Notes, "classic-six smoke path") {
				t.Fatalf("expected classic-six note for %s, got %v", tc.family, result.Notes)
			}
		})
	}
}

func TestAgentWorkflowsSmokeEvidence(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{
		"--dev-type", DevTypeFullStack.String(),
		"--team",
		"--complexity", ComplexityLow.String(),
		"--seed", "17",
		"--minimal",
		"--no-color",
	}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d: %s", code, stderr.String())
	}

	output := stdout.String()
	if !strings.Contains(output, FamilyAgentWorkflows.String()) {
		t.Fatalf("expected agent_workflows output, got %q", output)
	}
	if !strings.Contains(output, "agent_workflows depth pass") {
		t.Fatalf("expected dedicated depth message, got %q", output)
	}
	if !strings.Contains(output, "rust-stakeholder") || !strings.Contains(output, "stakeholder-core") {
		t.Fatalf("expected traceability markers, got %q", output)
	}
}

func TestModernCoreDedicatedRenderers(t *testing.T) {
	cases := []struct {
		name          string
		family        Family
		dev           DevType
		focusKey      string
		focusValue    string
		javaAnchor    string
		detailSnippet string
	}{
		{name: "agent_workflows", family: FamilyAgentWorkflows, dev: DevTypeFullStack, focusKey: "coordinationMode", focusValue: "delegated agent work, approval gates, and cross-repo handoff envelopes", javaAnchor: "src/main/java/com/stakeholder/generators/AgentWorkflowsRenderer.java", detailSnippet: "routing coding-agent work through review queues and approval gates"},
		{name: "platform_engineering", family: FamilyPlatformEngineering, dev: DevTypeDevOps, focusKey: "platformSurface", focusValue: "golden paths, identity boundaries, and queue ownership in the shared platform lane", javaAnchor: "src/main/java/com/stakeholder/generators/PlatformEngineeringRenderer.java", detailSnippet: "lining up golden paths, identity federation, queue ownership, and paved-road rollouts"},
		{name: "observability_ai_runtime", family: FamilyObservabilityAIRuntime, dev: DevTypeSystemsProgramming, focusKey: "runtimeSignals", focusValue: "trace spans, token burn, GPU pressure, and policy denials in one runtime lane", javaAnchor: "src/main/java/com/stakeholder/generators/ObservabilityAIRuntimeRenderer.java", detailSnippet: "correlating inference spans, token burn, GPU saturation, and sandbox denials"},
		{name: "delivery_preview_ops", family: FamilyDeliveryPreviewOps, dev: DevTypeGameDevelopment, focusKey: "deliveryGuardrail", focusValue: "preview deploys, canaries, release flags, and rollback checkpoints under seed control", javaAnchor: "src/main/java/com/stakeholder/generators/DeliveryPreviewOpsRenderer.java", detailSnippet: "coordinating preview deploys, canary health, release flags, and rollback checkpoints"},
		{name: "supply_chain_security", family: FamilySupplyChainSecurity, dev: DevTypeSecurity, focusKey: "supplyChainPosture", focusValue: "provenance, attestations, dependency drift, and secret exposure in one security lane", javaAnchor: "src/main/java/com/stakeholder/generators/SupplyChainSecurityRenderer.java", detailSnippet: "linking attestations, dependency drift, key rotation, and registry trust signals"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := renderFamily(RenderContext{
				Config: Config{DevType: tc.dev},
				Family: tc.family,
				Spec:   FamilySpecFor(tc.family),
			})
			if !strings.Contains(result.Message, tc.detailSnippet) {
				t.Fatalf("expected modern-core detail for %s, got %q", tc.family, result.Message)
			}
			if !strings.Contains(result.Message, "Traceability is anchored to Java, Rust, and stakeholder-core.") {
				t.Fatalf("expected traceability sentence for %s, got %q", tc.family, result.Message)
			}
			if !containsString(result.Evidence, "dedicated-depth") {
				t.Fatalf("expected dedicated-depth evidence for %s, got %v", tc.family, result.Evidence)
			}
			if !containsString(result.Evidence, "java-stakeholder") {
				t.Fatalf("expected java-stakeholder evidence for %s, got %v", tc.family, result.Evidence)
			}
			if !containsString(result.Notes, "modern-core dedicated path") {
				t.Fatalf("expected modern-core note for %s, got %v", tc.family, result.Notes)
			}
			if !containsString(result.Notes, "focusKey="+tc.focusKey) {
				t.Fatalf("expected focus key note for %s, got %v", tc.family, result.Notes)
			}
			if !containsString(result.Notes, tc.focusKey+"="+tc.focusValue) {
				t.Fatalf("expected focus value note for %s, got %v", tc.family, result.Notes)
			}
			if !containsString(result.Notes, "javaPath="+tc.javaAnchor) {
				t.Fatalf("expected java anchor note for %s, got %v", tc.family, result.Notes)
			}
		})
	}
}

func TestTextModeUsesANSIStyling(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{
		"--dev-type", DevTypeBackend.String(),
		"--complexity", ComplexityLow.String(),
		"--seed", "11",
	}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d: %s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "\x1b[") {
		t.Fatalf("expected ANSI styling in text output, got %q", stdout.String())
	}
}

func TestExperimentalFlagsFailFast(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{
		"--experimental-provider", "openai-compatible",
		"--experimental-model", "gpt-5.4",
		"--experimental-profile", "consumer-openai",
		"--experimental-prompt", "tailored",
		"--experimental-adapter-mode", ExperimentalAdapterModeConsumer.String(),
	}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("expected exit 2, got %d", code)
	}
	if !strings.Contains(stderr.String(), "experimental provider mode is not implemented in go-stakeholder") {
		t.Fatalf("expected explicit not-implemented error, got %q", stderr.String())
	}
	if !strings.Contains(stderr.String(), "adapter-mode=consumer") {
		t.Fatalf("expected adapter-mode detail, got %q", stderr.String())
	}
}

func TestGroupedFallbackRenderer(t *testing.T) {
	result := renderFamily(RenderContext{
		Config: Config{DevType: DevTypeDevOps},
		Family: FamilyAIInferenceOps,
		Spec:   FamilySpecFor(FamilyAIInferenceOps),
	})
	if !strings.Contains(result.Message, "ai-governance fallback") {
		t.Fatalf("expected grouped fallback message, got %q", result.Message)
	}
	if !containsString(result.Evidence, "grouped-fallback") {
		t.Fatalf("expected grouped fallback evidence, got %v", result.Evidence)
	}
}

func decodeEvents(t *testing.T, raw []byte) []SessionEvent {
	t.Helper()

	scanner := bufio.NewScanner(bytes.NewReader(raw))
	events := make([]SessionEvent, 0, 4)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var event SessionEvent
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			t.Fatalf("decode event %q: %v", line, err)
		}
		events = append(events, event)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scan events: %v", err)
	}
	return events
}

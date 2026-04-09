package app

import (
	"fmt"
	"strings"
)

type rendererFunc func(RenderContext) RenderResult

var dedicatedRenderers = map[Family]rendererFunc{
	FamilyAgentWorkflows:   renderAgentWorkflows,
	FamilyCodeAnalyzer:     renderCodeAnalyzer,
	FamilyDataProcessing:   renderDataProcessing,
	FamilyJargon:           renderJargon,
	FamilyMetrics:          renderMetrics,
	FamilyNetworkActivity:  renderNetworkActivity,
	FamilySystemMonitoring: renderSystemMonitoring,
}

var groupedRenderers = map[FamilyGroup]rendererFunc{
	FamilyGroupClassicSix:         renderClassicSixFallback,
	FamilyGroupModernCore:         renderModernCoreFallback,
	FamilyGroupAIGovernance:       renderAIGovernanceFallback,
	FamilyGroupSecurityBlockchain: renderSecurityBlockchainFallback,
	FamilyGroupHealthProtocol:     renderHealthProtocolFallback,
	FamilyGroupOverlayQuantum:     renderOverlayQuantumFallback,
}

func renderFamily(ctx RenderContext) RenderResult {
	if fn, ok := dedicatedRenderers[ctx.Family]; ok {
		return fn(ctx)
	}
	if fn, ok := groupedRenderers[ctx.Spec.Group]; ok {
		return fn(ctx)
	}
	return renderClassicSixFallback(ctx)
}

func renderCodeAnalyzer(ctx RenderContext) RenderResult {
	return smokeResult(ctx, "code_analyzer", "classic-six smoke path")
}

func renderDataProcessing(ctx RenderContext) RenderResult {
	return smokeResult(ctx, "data_processing", "classic-six smoke path")
}

func renderJargon(ctx RenderContext) RenderResult {
	return smokeResult(ctx, "jargon", "classic-six smoke path")
}

func renderMetrics(ctx RenderContext) RenderResult {
	return smokeResult(ctx, "metrics", "classic-six smoke path")
}

func renderNetworkActivity(ctx RenderContext) RenderResult {
	return smokeResult(ctx, "network_activity", "classic-six smoke path")
}

func renderSystemMonitoring(ctx RenderContext) RenderResult {
	return smokeResult(ctx, "system_monitoring", "classic-six smoke path")
}

func renderAgentWorkflows(ctx RenderContext) RenderResult {
	project := ctx.Config.Project
	if project == "" {
		project = "stakeholder"
	}
	return RenderResult{
		Message: fmt.Sprintf("agent_workflows smoke evidence: %s handoff loop completed under %d seed", project, ctx.Config.Seed),
		Evidence: []string{
			"dedicated-smoke",
			"traceability-linked",
			"rust-stakeholder",
			"stakeholder-core",
		},
		Notes: []string{
			"family-specific renderer",
			"modern-core smoke path",
		},
	}
}

func smokeResult(ctx RenderContext, family string, note string) RenderResult {
	project := ctx.Config.Project
	if project == "" {
		project = "stakeholder"
	}
	return RenderResult{
		Message: fmt.Sprintf("%s smoke evidence: %s kept the audited contract aligned for %s", family, project, ctx.Config.DevType.String()),
		Evidence: []string{
			"dedicated-smoke",
			"traceability-linked",
			"rust-stakeholder",
			"stakeholder-core",
		},
		Notes: []string{
			"family-specific renderer",
			note,
		},
	}
}

func renderClassicSixFallback(ctx RenderContext) RenderResult {
	return groupedFallbackResult(ctx, "classic-six", "baseline contract sweep")
}

func renderModernCoreFallback(ctx RenderContext) RenderResult {
	return groupedFallbackResult(ctx, "modern-core", "workflow and platform sweep")
}

func renderAIGovernanceFallback(ctx RenderContext) RenderResult {
	return groupedFallbackResult(ctx, "ai-governance", "model and retrieval sweep")
}

func renderSecurityBlockchainFallback(ctx RenderContext) RenderResult {
	return groupedFallbackResult(ctx, "security-blockchain", "security and chain sweep")
}

func renderHealthProtocolFallback(ctx RenderContext) RenderResult {
	return groupedFallbackResult(ctx, "health-protocol", "protocol interoperability sweep")
}

func renderOverlayQuantumFallback(ctx RenderContext) RenderResult {
	return groupedFallbackResult(ctx, "overlay-quantum", "overlay and quantum sweep")
}

func groupedFallbackResult(ctx RenderContext, label string, description string) RenderResult {
	message := fmt.Sprintf("%s fallback: %s handled by %s renderer for %s", label, ctx.Family.String(), ctx.Spec.Renderer, description)
	if ctx.Config.Minimal {
		message = fmt.Sprintf("%s fallback: %s", label, ctx.Family.String())
	}
	return RenderResult{
		Message: message,
		Evidence: []string{
			"grouped-fallback",
			"traceability-linked",
		},
		Notes: []string{
			fmt.Sprintf("group=%s", label),
			fmt.Sprintf("family=%s", ctx.Family.String()),
		},
	}
}

func formatTextEvent(event SessionEvent, noColor bool) string {
	head := styleForEventType(event.EventType, noColor)
	body := event.Message
	if event.Context.Activity != nil {
		activity := event.Context.Activity
		activityLabel := activity.Family
		groupLabel := activity.Group
		rendererLabel := activity.Renderer
		if !noColor {
			activityLabel = ansi("1;32", activityLabel)
			groupLabel = ansi("1;35", groupLabel)
			rendererLabel = ansi("0;36", rendererLabel)
		}
		body = fmt.Sprintf("%s family=%s group=%s renderer=%s", body, activityLabel, groupLabel, rendererLabel)
	}
	if len(event.Context.Traceability) > 0 {
		headTrace := event.Context.Traceability[0]
		traceLabel := fmt.Sprintf("%s->%s", headTrace.Source, headTrace.Target)
		if !noColor {
			traceLabel = ansi("0;90", traceLabel)
		}
		body = fmt.Sprintf("%s trace=%s", body, traceLabel)
	}
	return fmt.Sprintf("%s %s", head, body)
}

func styleForEventType(eventType string, noColor bool) string {
	if noColor {
		return fmt.Sprintf("[%s]", eventType)
	}
	switch eventType {
	case "session.start":
		return ansi("1;36", fmt.Sprintf("[%s]", eventType))
	case "activity":
		return ansi("1;32", fmt.Sprintf("[%s]", eventType))
	case "session.end":
		return ansi("1;33", fmt.Sprintf("[%s]", eventType))
	default:
		return ansi("1;35", fmt.Sprintf("[%s]", eventType))
	}
}

func ansi(code string, value string) string {
	return "\x1b[" + code + "m" + value + "\x1b[0m"
}

func formatTextStream(events []SessionEvent, minimal bool, noColor bool) []string {
	lines := make([]string, 0, len(events))
	for _, event := range events {
		if minimal && event.EventType != "activity" {
			continue
		}
		lines = append(lines, formatTextEvent(event, noColor))
	}
	return lines
}

func hasExperimentalInput(cfg ExperimentalConfig) bool {
	return cfg.Provider != "" || cfg.Model != "" || cfg.Profile != "" || cfg.Prompt != "" || cfg.AdapterMode != ""
}

func validateExperimentalConfig(cfg ExperimentalConfig) error {
	if cfg.AdapterMode != "" && cfg.AdapterMode != ExperimentalAdapterModeAPI && cfg.AdapterMode != ExperimentalAdapterModeConsumer {
		return fmt.Errorf("invalid experimental adapter mode %q", cfg.AdapterMode)
	}
	if cfg.Provider == "" && cfg.Model == "" && cfg.Profile == "" && cfg.Prompt == "" && cfg.AdapterMode == "" {
		return nil
	}
	allowed := map[string]struct{}{
		"anthropic":         {},
		"claude-consumer":   {},
		"openai-compatible": {},
		"openai-consumer":   {},
	}
	if cfg.Provider != "" {
		if _, ok := allowed[strings.ToLower(cfg.Provider)]; !ok {
			return fmt.Errorf("invalid experimental provider %q", cfg.Provider)
		}
	}
	return nil
}

func normalizedExperimentalConfig(raw ExperimentalConfig) ExperimentalConfig {
	return ExperimentalConfig{
		Provider:    strings.ToLower(strings.TrimSpace(raw.Provider)),
		Model:       strings.TrimSpace(raw.Model),
		Profile:     strings.TrimSpace(raw.Profile),
		Prompt:      strings.TrimSpace(raw.Prompt),
		AdapterMode: ExperimentalAdapterMode(strings.ToLower(strings.TrimSpace(string(raw.AdapterMode)))),
		Enabled:     hasExperimentalInput(raw),
	}
}

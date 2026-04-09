package app

import (
	"math/rand"
	"sort"
	"time"
)

var sessionBaseTime = time.Date(2026, time.April, 9, 0, 0, 0, 0, time.UTC)

func BuildSession(cfg Config) []SessionEvent {
	selections := SelectFamilies(cfg)
	sessionContext := SessionContext{
		DevType:          cfg.DevType.String(),
		Jargon:           cfg.Jargon.String(),
		Complexity:       cfg.Complexity.String(),
		Duration:         cfg.Duration,
		Alerts:           cfg.Alerts,
		Project:          cfg.Project,
		Minimal:          cfg.Minimal,
		Team:             cfg.Team,
		Framework:        cfg.Framework,
		Seed:             cfg.Seed,
		OutputFormat:     cfg.OutputFormat.String(),
		NoColor:          cfg.NoColor,
		Trace:            cfg.Trace,
		SelectedFamilies: selectedFamilyNames(selections),
	}

	events := make([]SessionEvent, 0, len(selections)+2)
	events = append(events, SessionEvent{
		EventType: "session.start",
		Sequence:  1,
		Message:   "deterministic session started",
		Timestamp: eventTimestamp(1),
		Context: EventContext{
			Session: sessionContext,
			Notes:   []string{"shared engine", "registry-based dispatch"},
		},
	})

	for i, selection := range selections {
		spec := FamilySpecFor(selection.Family)
		rendered := renderFamily(RenderContext{
			Config:           cfg,
			Family:           selection.Family,
			Spec:             spec,
			SelectionIndex:   i,
			TotalSelections:  len(selections),
			SelectedFamilies: selections,
		})
		events = append(events, SessionEvent{
			EventType: "activity",
			Sequence:  len(events) + 1,
			Message:   rendered.Message,
			Timestamp: eventTimestamp(len(events) + 1),
			Context: EventContext{
				Session: sessionContext,
				Activity: &ActivityContext{
					Family:   selection.Family.String(),
					Group:    spec.Group.String(),
					Renderer: spec.Renderer,
					RuleTags: append([]string(nil), selection.RuleTags...),
					Smoke:    spec.Smoke,
				},
				Traceability: traceabilityFor(spec, selection),
				Evidence:     rendered.Evidence,
				Notes:        rendered.Notes,
			},
		})
	}

	events = append(events, SessionEvent{
		EventType: "session.end",
		Sequence:  len(events) + 1,
		Message:   "deterministic session finished",
		Timestamp: eventTimestamp(len(events) + 1),
		Context: EventContext{
			Session: sessionContext,
			Notes:   []string{"normalized-json", "ansi-text", "smoke-evidence"},
		},
	})

	return events
}

func SelectFamilies(cfg Config) []SelectedFamily {
	base := primaryFamilyForDevType(cfg.DevType)
	required := []SelectedFamily{{Family: base, RuleTags: []string{"base"}}}

	if cfg.Complexity == ComplexityMedium || cfg.Complexity == ComplexityHigh || cfg.Complexity == ComplexityExtreme {
		modern := modernFamilyForDevType(cfg.DevType)
		if modern != "" {
			required = append(required, SelectedFamily{Family: modern, RuleTags: []string{"modern"}})
		}
	}

	if cfg.Complexity == ComplexityHigh || cfg.Complexity == ComplexityExtreme {
		policy := policyFamilyForDevType(cfg.DevType)
		if policy != "" {
			required = append(required, SelectedFamily{Family: policy, RuleTags: []string{"policy"}})
		}
	}

	if cfg.Team {
		team := teamFamilyForDevType(cfg.DevType)
		if team != "" {
			required = append(required, SelectedFamily{Family: team, RuleTags: []string{"team"}})
		}
	}

	if cfg.Alerts {
		alert := alertFamilyForDevType(cfg.DevType)
		if alert != "" {
			required = append(required, SelectedFamily{Family: alert, RuleTags: []string{"alerts"}})
		}
	}

	selectionMap := make(map[Family]SelectedFamily, len(required))
	selected := make([]SelectedFamily, 0, len(required))
	for _, item := range required {
		if _, ok := selectionMap[item.Family]; ok {
			continue
		}
		selectionMap[item.Family] = item
		selected = append(selected, item)
	}

	targetCount := cfg.Complexity.ActivityCount()
	if len(selected) > targetCount {
		targetCount = len(selected)
	}

	rng := rand.New(rand.NewSource(cfg.Seed))
	pool := seededFallbackPool(cfg.DevType)
	for len(selected) < targetCount && len(pool) > 0 {
		nextIndex := int(rng.Int63n(int64(len(pool))))
		candidate := pool[nextIndex]
		pool = append(pool[:nextIndex], pool[nextIndex+1:]...)
		if containsSelected(selected, candidate) {
			continue
		}
		selected = append(selected, SelectedFamily{Family: candidate, RuleTags: []string{"filler"}})
	}

	return selected
}

func containsSelected(list []SelectedFamily, family Family) bool {
	for _, item := range list {
		if item.Family == family {
			return true
		}
	}
	return false
}

func selectedFamilyNames(list []SelectedFamily) []string {
	out := make([]string, 0, len(list))
	for _, item := range list {
		out = append(out, item.Family.String())
	}
	return out
}

func eventTimestamp(sequence int) string {
	return sessionBaseTime.Add(time.Duration(sequence) * time.Second).Format(time.RFC3339)
}

func traceabilityFor(spec FamilySpec, selection SelectedFamily) []TraceRef {
	base := []TraceRef{
		{
			Source: "rust-stakeholder",
			Target: "stakeholder-core",
			Kind:   "baseline",
			Note:   "shared deterministic family registry",
		},
		{
			Source: "stakeholder-core",
			Target: "go-stakeholder",
			Kind:   "packet",
			Note:   "follower foundation and smoke evidence",
		},
	}
	if spec.Smoke {
		base = append(base, TraceRef{
			Source: "stakeholder-core",
			Target: spec.Name.String(),
			Kind:   "smoke",
			Note:   "dedicated smoke evidence",
		})
	}
	return base
}

func primaryFamilyForDevType(dev DevType) Family {
	switch dev {
	case DevTypeBackend:
		return FamilyCodeAnalyzer
	case DevTypeBlockchain:
		return FamilyBlockchainProtocolOps
	case DevTypeDataScience:
		return FamilyDataProcessing
	case DevTypeDevOps:
		return FamilySystemMonitoring
	case DevTypeFrontend:
		return FamilyEdgeClientRuntime
	case DevTypeGameDevelopment:
		return FamilyDeliveryPreviewOps
	case DevTypeMachineLearning:
		return FamilyAIInferenceOps
	case DevTypeSecurity:
		return FamilySupplyChainSecurity
	case DevTypeSystemsProgramming:
		return FamilyCompilerMaintainer
	case DevTypeFullStack:
		return FamilyCodeAnalyzer
	default:
		return FamilyCodeAnalyzer
	}
}

func modernFamilyForDevType(dev DevType) Family {
	switch dev {
	case DevTypeBackend, DevTypeFullStack:
		return FamilyAgentWorkflows
	case DevTypeBlockchain:
		return FamilyCrossChainInterop
	case DevTypeDataScience:
		return FamilyAIInferenceOps
	case DevTypeDevOps, DevTypeSystemsProgramming:
		return FamilyObservabilityAIRuntime
	case DevTypeFrontend, DevTypeGameDevelopment:
		return FamilyDeliveryPreviewOps
	case DevTypeMachineLearning:
		return FamilyKnowledgeRetrieval
	case DevTypeSecurity:
		return FamilyAgentBoundarySecurity
	default:
		return FamilyPlatformEngineering
	}
}

func policyFamilyForDevType(dev DevType) Family {
	switch dev {
	case DevTypeBackend, DevTypeFullStack:
		return FamilySupplyChainSecurity
	case DevTypeBlockchain:
		return FamilyIdentityAndTrust
	case DevTypeDataScience, DevTypeMachineLearning:
		return FamilyAIBOMProvenance
	case DevTypeDevOps, DevTypeSystemsProgramming:
		return FamilyObservabilityAIRuntime
	case DevTypeFrontend, DevTypeGameDevelopment:
		return FamilyEdgeClientRuntime
	case DevTypeSecurity:
		return FamilyAgentBoundarySecurity
	default:
		return FamilySupplyChainSecurity
	}
}

func teamFamilyForDevType(dev DevType) Family {
	switch dev {
	case DevTypeBlockchain:
		return FamilyMCPA2AOps
	case DevTypeDataScience, DevTypeMachineLearning:
		return FamilyAgentWorkflows
	case DevTypeDevOps, DevTypeSystemsProgramming:
		return FamilyPlatformEngineering
	case DevTypeFrontend, DevTypeGameDevelopment:
		return FamilyDeliveryPreviewOps
	case DevTypeSecurity:
		return FamilyAgentWorkflows
	default:
		return FamilyAgentWorkflows
	}
}

func alertFamilyForDevType(dev DevType) Family {
	switch dev {
	case DevTypeBlockchain:
		return FamilyProofAndSequencerOps
	case DevTypeDataScience, DevTypeMachineLearning:
		return FamilyObservabilityAIRuntime
	case DevTypeDevOps, DevTypeSystemsProgramming:
		return FamilySupplyChainSecurity
	case DevTypeFrontend, DevTypeGameDevelopment:
		return FamilyStreamingBusOps
	case DevTypeSecurity:
		return FamilySupplyChainSecurity
	default:
		return FamilyObservabilityAIRuntime
	}
}

func seededFallbackPool(dev DevType) []Family {
	ordered := []Family{
		FamilyCodeAnalyzer,
		FamilyDataProcessing,
		FamilyJargon,
		FamilyMetrics,
		FamilyNetworkActivity,
		FamilySystemMonitoring,
		FamilyAgentWorkflows,
		FamilyAIInferenceOps,
		FamilyPlatformEngineering,
		FamilySupplyChainSecurity,
		FamilyObservabilityAIRuntime,
		FamilyDeliveryPreviewOps,
		FamilyEvaluationAndGuardrails,
		FamilyKnowledgeRetrieval,
		FamilyEdgeClientRuntime,
		FamilyIdentityAndTrust,
		FamilyAIBOMProvenance,
		FamilyAgentBoundarySecurity,
		FamilyEmbeddedAgenticPipeline,
		FamilyDataGovernanceCompliance,
		FamilyFinOpsCapacity,
		FamilyBlockchainProtocolOps,
		FamilyCrossChainInterop,
		FamilyProofAndSequencerOps,
		FamilyHybridRuntimeOps,
		FamilyCapacityCostController,
		FamilyBatchExecutionTuner,
		FamilyCompilerMaintainer,
		FamilyInteropAdapterEngineer,
		FamilyPreflightCapacityPlanner,
		FamilySimulatorPerformanceEngineer,
		FamilyFHIRProfileGenerator,
		FamilySMARTLaunchOAuth,
		FamilyBulkFHIRPopulationOps,
		FamilyHL7V2FeedOps,
		FamilyClinicalWorkflowEvents,
		FamilyDICOMWebImagingOps,
		FamilyOpenEHRSemanticRecordOps,
		FamilyDeviceTelemetryClinical,
		FamilyEMRVendorAdapter,
		FamilyOCPPChargePointOps,
		FamilyOCPIRoamingOps,
		FamilyMCPA2AOps,
		FamilyStreamingBusOps,
		FamilyServiceMeshRpcOps,
		FamilyMultilingualSecurityPacks,
		FamilySecurityPersonaPacks,
	}

	excluded := map[Family]struct{}{
		primaryFamilyForDevType(dev): {},
	}
	if mod := modernFamilyForDevType(dev); mod != "" {
		excluded[mod] = struct{}{}
	}
	if pol := policyFamilyForDevType(dev); pol != "" {
		excluded[pol] = struct{}{}
	}
	if team := teamFamilyForDevType(dev); team != "" {
		excluded[team] = struct{}{}
	}
	if alert := alertFamilyForDevType(dev); alert != "" {
		excluded[alert] = struct{}{}
	}

	pool := make([]Family, 0, len(ordered))
	for _, candidate := range ordered {
		if _, ok := excluded[candidate]; ok {
			continue
		}
		pool = append(pool, candidate)
	}

	sort.SliceStable(pool, func(i, j int) bool { return pool[i] < pool[j] })
	return pool
}

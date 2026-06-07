package app

import "strings"

var familyRegistry = []FamilySpec{
	{Name: FamilyCodeAnalyzer, Group: FamilyGroupClassicSix, Smoke: true, Renderer: "dedicated:code_analyzer"},
	{Name: FamilyDataProcessing, Group: FamilyGroupClassicSix, Smoke: true, Renderer: "dedicated:data_processing"},
	{Name: FamilyJargon, Group: FamilyGroupClassicSix, Smoke: true, Renderer: "dedicated:jargon"},
	{Name: FamilyMetrics, Group: FamilyGroupClassicSix, Smoke: true, Renderer: "dedicated:metrics"},
	{Name: FamilyNetworkActivity, Group: FamilyGroupClassicSix, Smoke: true, Renderer: "dedicated:network_activity"},
	{Name: FamilySystemMonitoring, Group: FamilyGroupClassicSix, Smoke: true, Renderer: "dedicated:system_monitoring"},
	{Name: FamilyAgentWorkflows, Group: FamilyGroupModernCore, Smoke: true, Renderer: "dedicated:agent_workflows"},
	{Name: FamilyAIInferenceOps, Group: FamilyGroupAIGovernance, Smoke: false, Renderer: "group:ai-governance"},
	{Name: FamilyPlatformEngineering, Group: FamilyGroupModernCore, Smoke: true, Renderer: "dedicated:platform_engineering"},
	{Name: FamilySupplyChainSecurity, Group: FamilyGroupModernCore, Smoke: true, Renderer: "dedicated:supply_chain_security"},
	{Name: FamilyObservabilityAIRuntime, Group: FamilyGroupModernCore, Smoke: true, Renderer: "dedicated:observability_ai_runtime"},
	{Name: FamilyDeliveryPreviewOps, Group: FamilyGroupModernCore, Smoke: true, Renderer: "dedicated:delivery_preview_ops"},
	{Name: FamilyEvaluationAndGuardrails, Group: FamilyGroupAIGovernance, Smoke: false, Renderer: "group:ai-governance"},
	{Name: FamilyKnowledgeRetrieval, Group: FamilyGroupAIGovernance, Smoke: false, Renderer: "group:ai-governance"},
	{Name: FamilyEdgeClientRuntime, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyIdentityAndTrust, Group: FamilyGroupSecurityBlockchain, Smoke: false, Renderer: "group:security-blockchain"},
	{Name: FamilyAIBOMProvenance, Group: FamilyGroupAIGovernance, Smoke: false, Renderer: "group:ai-governance"},
	{Name: FamilyAgentBoundarySecurity, Group: FamilyGroupSecurityBlockchain, Smoke: false, Renderer: "group:security-blockchain"},
	{Name: FamilyEmbeddedAgenticPipeline, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyDataGovernanceCompliance, Group: FamilyGroupAIGovernance, Smoke: false, Renderer: "group:ai-governance"},
	{Name: FamilyFinOpsCapacity, Group: FamilyGroupAIGovernance, Smoke: false, Renderer: "group:ai-governance"},
	{Name: FamilyBlockchainProtocolOps, Group: FamilyGroupSecurityBlockchain, Smoke: false, Renderer: "group:security-blockchain"},
	{Name: FamilyCrossChainInterop, Group: FamilyGroupSecurityBlockchain, Smoke: false, Renderer: "group:security-blockchain"},
	{Name: FamilyProofAndSequencerOps, Group: FamilyGroupSecurityBlockchain, Smoke: false, Renderer: "group:security-blockchain"},
	{Name: FamilyHybridRuntimeOps, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilyCapacityCostController, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilyBatchExecutionTuner, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilyCompilerMaintainer, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilyInteropAdapterEngineer, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilyPreflightCapacityPlanner, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilySimulatorPerformanceEngineer, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilyFHIRProfileGenerator, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilySMARTLaunchOAuth, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyBulkFHIRPopulationOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyHL7V2FeedOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyClinicalWorkflowEvents, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyDICOMWebImagingOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyOpenEHRSemanticRecordOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyDeviceTelemetryClinical, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyEMRVendorAdapter, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyOCPPChargePointOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyOCPIRoamingOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyMCPA2AOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyStreamingBusOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyServiceMeshRpcOps, Group: FamilyGroupHealthProtocol, Smoke: false, Renderer: "group:health-protocol"},
	{Name: FamilyMultilingualSecurityPacks, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
	{Name: FamilySecurityPersonaPacks, Group: FamilyGroupOverlayQuantum, Smoke: false, Renderer: "group:overlay-quantum"},
}

var familyRegistryByName = func() map[Family]FamilySpec {
	out := make(map[Family]FamilySpec, len(familyRegistry))
	for _, spec := range familyRegistry {
		out[spec.Name] = spec
	}
	return out
}()

func FamilyRegistry() []FamilySpec {
	out := make([]FamilySpec, len(familyRegistry))
	copy(out, familyRegistry)
	return out
}

func FamilySpecFor(name Family) FamilySpec {
	if spec, ok := familyRegistryByName[name]; ok {
		return spec
	}
	return FamilySpec{Name: name, Group: FamilyGroupClassicSix, Renderer: "group:classic-six"}
}

func valueCatalog() ValueCatalog {
	var catalog ValueCatalog
	catalog.Config.DevType = devTypeValues()
	catalog.Config.Jargon = jargonValues()
	catalog.Config.Complexity = complexityValues()
	catalog.Config.OutputFormat = outputFormatValues()
	catalog.Flags.Alerts = []bool{false, true}
	catalog.Flags.Team = []bool{false, true}
	catalog.Flags.Minimal = []bool{false, true}
	catalog.Flags.NoColor = []bool{false, true}
	catalog.Flags.Trace = []bool{false, true}
	catalog.Experimental.Provider = experimentalProviders()
	catalog.Experimental.AdapterMode = []string{
		ExperimentalAdapterModeAPI.String(),
		ExperimentalAdapterModeConsumer.String(),
	}
	catalog.Families = FamilyRegistry()
	return catalog
}

func devTypeValues() []string {
	return []string{
		DevTypeBackend.String(),
		DevTypeBlockchain.String(),
		DevTypeDataScience.String(),
		DevTypeDevOps.String(),
		DevTypeFrontend.String(),
		DevTypeFullStack.String(),
		DevTypeGameDevelopment.String(),
		DevTypeMachineLearning.String(),
		DevTypeSecurity.String(),
		DevTypeSystemsProgramming.String(),
	}
}

func jargonValues() []string {
	return []string{
		JargonLevelLow.String(),
		JargonLevelMedium.String(),
		JargonLevelHigh.String(),
		JargonLevelExtreme.String(),
	}
}

func complexityValues() []string {
	return []string{
		ComplexityLow.String(),
		ComplexityMedium.String(),
		ComplexityHigh.String(),
		ComplexityExtreme.String(),
	}
}

func outputFormatValues() []string {
	return []string{
		OutputFormatText.String(),
		OutputFormatJSON.String(),
	}
}

func experimentalProviders() []string {
	return []string{
		"anthropic",
		"claude-consumer",
		"openai-compatible",
		"openai-consumer",
	}
}

func registryFamilyNames() []Family {
	out := make([]Family, len(familyRegistry))
	for i, spec := range familyRegistry {
		out[i] = spec.Name
	}
	return out
}

func containsFamily(list []Family, needle Family) bool {
	for _, item := range list {
		if item == needle {
			return true
		}
	}
	return false
}

func containsString(list []string, needle string) bool {
	for _, item := range list {
		if item == needle {
			return true
		}
	}
	return false
}

func groupLabel(group FamilyGroup) string {
	return strings.ReplaceAll(string(group), "-", " ")
}

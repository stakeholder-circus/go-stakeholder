package app

import "fmt"

type DevType string

const (
	DevTypeBackend            DevType = "backend"
	DevTypeDataScience        DevType = "data-science"
	DevTypeDevOps             DevType = "dev-ops"
	DevTypeFrontend           DevType = "frontend"
	DevTypeFullStack          DevType = "fullstack"
	DevTypeGameDevelopment    DevType = "game-development"
	DevTypeMachineLearning    DevType = "machine-learning"
	DevTypeSecurity           DevType = "security"
	DevTypeSystemsProgramming DevType = "systems-programming"
	DevTypeBlockchain         DevType = "blockchain"
)

type JargonLevel string

const (
	JargonLevelLow     JargonLevel = "low"
	JargonLevelMedium  JargonLevel = "medium"
	JargonLevelHigh    JargonLevel = "high"
	JargonLevelExtreme JargonLevel = "extreme"
)

type Complexity string

const (
	ComplexityLow     Complexity = "low"
	ComplexityMedium  Complexity = "medium"
	ComplexityHigh    Complexity = "high"
	ComplexityExtreme Complexity = "extreme"
)

func (c Complexity) ActivityCount() int {
	switch c {
	case ComplexityLow:
		return 1
	case ComplexityMedium:
		return 2
	case ComplexityHigh:
		return 3
	case ComplexityExtreme:
		return 4
	default:
		return 2
	}
}

type OutputFormat string

const (
	OutputFormatText OutputFormat = "text"
	OutputFormatJSON OutputFormat = "json"
)

type ExperimentalAdapterMode string

const (
	ExperimentalAdapterModeAPI      ExperimentalAdapterMode = "api"
	ExperimentalAdapterModeConsumer ExperimentalAdapterMode = "consumer"
)

type Family string

const (
	FamilyCodeAnalyzer                 Family = "code_analyzer"
	FamilyDataProcessing               Family = "data_processing"
	FamilyJargon                       Family = "jargon"
	FamilyMetrics                      Family = "metrics"
	FamilyNetworkActivity              Family = "network_activity"
	FamilySystemMonitoring             Family = "system_monitoring"
	FamilyAgentWorkflows               Family = "agent_workflows"
	FamilyAIInferenceOps               Family = "ai_inference_ops"
	FamilyPlatformEngineering          Family = "platform_engineering"
	FamilySupplyChainSecurity          Family = "supply_chain_security"
	FamilyObservabilityAIRuntime       Family = "observability_ai_runtime"
	FamilyDeliveryPreviewOps           Family = "delivery_preview_ops"
	FamilyEvaluationAndGuardrails      Family = "evaluation_and_guardrails"
	FamilyKnowledgeRetrieval           Family = "knowledge_retrieval"
	FamilyEdgeClientRuntime            Family = "edge_client_runtime"
	FamilyIdentityAndTrust             Family = "identity_and_trust"
	FamilyAIBOMProvenance              Family = "aibom_provenance"
	FamilyAgentBoundarySecurity        Family = "agent_boundary_security"
	FamilyEmbeddedAgenticPipeline      Family = "embedded_agentic_pipeline"
	FamilyDataGovernanceCompliance     Family = "data_governance_compliance"
	FamilyFinOpsCapacity               Family = "finops_capacity"
	FamilyBlockchainProtocolOps        Family = "blockchain_protocol_ops"
	FamilyCrossChainInterop            Family = "cross_chain_interop"
	FamilyProofAndSequencerOps         Family = "proof_and_sequencer_ops"
	FamilyHybridRuntimeOps             Family = "hybrid_runtime_ops"
	FamilyCapacityCostController       Family = "capacity_cost_controller"
	FamilyBatchExecutionTuner          Family = "batch_execution_tuner"
	FamilyCompilerMaintainer           Family = "compiler_maintainer"
	FamilyInteropAdapterEngineer       Family = "interop_adapter_engineer"
	FamilyPreflightCapacityPlanner     Family = "preflight_capacity_planner"
	FamilySimulatorPerformanceEngineer Family = "simulator_performance_engineer"
	FamilyFHIRProfileGenerator         Family = "fhir_profile_generator"
	FamilySMARTLaunchOAuth             Family = "smart_launch_oauth"
	FamilyBulkFHIRPopulationOps        Family = "bulk_fhir_population_ops"
	FamilyHL7V2FeedOps                 Family = "hl7v2_feed_ops"
	FamilyClinicalWorkflowEvents       Family = "clinical_workflow_events"
	FamilyDICOMWebImagingOps           Family = "dicomweb_imaging_ops"
	FamilyOpenEHRSemanticRecordOps     Family = "openehr_semantic_record_ops"
	FamilyDeviceTelemetryClinical      Family = "device_telemetry_clinical"
	FamilyEMRVendorAdapter             Family = "emr_vendor_adapter"
	FamilyOCPPChargePointOps           Family = "ocpp_chargepoint_ops"
	FamilyOCPIRoamingOps               Family = "ocpi_roaming_ops"
	FamilyMCPA2AOps                    Family = "mcp_a2a_ops"
	FamilyStreamingBusOps              Family = "streaming_bus_ops"
	FamilyServiceMeshRpcOps            Family = "service_mesh_rpc_ops"
	FamilyMultilingualSecurityPacks    Family = "multilingual_security_packs"
	FamilySecurityPersonaPacks         Family = "security_persona_packs"
)

type FamilyGroup string

const (
	FamilyGroupClassicSix         FamilyGroup = "classic-six"
	FamilyGroupModernCore         FamilyGroup = "modern-core"
	FamilyGroupAIGovernance       FamilyGroup = "ai-governance"
	FamilyGroupSecurityBlockchain FamilyGroup = "security-blockchain"
	FamilyGroupHealthProtocol     FamilyGroup = "health-protocol"
	FamilyGroupOverlayQuantum     FamilyGroup = "overlay-quantum"
)

type FamilySpec struct {
	Name     Family      `json:"name"`
	Group    FamilyGroup `json:"group"`
	Smoke    bool        `json:"smoke"`
	Renderer string      `json:"renderer"`
}

type TraceRef struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Kind   string `json:"kind"`
	Note   string `json:"note,omitempty"`
}

type SessionContext struct {
	DevType          string   `json:"devType"`
	Jargon           string   `json:"jargon"`
	Complexity       string   `json:"complexity"`
	Duration         int      `json:"duration"`
	Alerts           bool     `json:"alerts"`
	Project          string   `json:"project"`
	Minimal          bool     `json:"minimal"`
	Team             bool     `json:"team"`
	Framework        string   `json:"framework"`
	Seed             int64    `json:"seed"`
	OutputFormat     string   `json:"outputFormat"`
	NoColor          bool     `json:"noColor"`
	Trace            bool     `json:"trace"`
	SelectedFamilies []string `json:"selectedFamilies"`
}

type ActivityContext struct {
	Family   string   `json:"family"`
	Group    string   `json:"group"`
	Renderer string   `json:"renderer"`
	RuleTags []string `json:"ruleTags,omitempty"`
	Smoke    bool     `json:"smoke"`
}

type EventContext struct {
	Session      SessionContext   `json:"session"`
	Activity     *ActivityContext `json:"activity,omitempty"`
	Traceability []TraceRef       `json:"traceability,omitempty"`
	Evidence     []string         `json:"evidence,omitempty"`
	Notes        []string         `json:"notes,omitempty"`
}

type SessionEvent struct {
	EventType string       `json:"eventType"`
	Sequence  int          `json:"sequence"`
	Message   string       `json:"message"`
	Timestamp string       `json:"timestamp"`
	Context   EventContext `json:"context"`
}

type SelectedFamily struct {
	Family   Family
	RuleTags []string
}

type RenderContext struct {
	Config           Config
	Family           Family
	Spec             FamilySpec
	SelectionIndex   int
	TotalSelections  int
	SelectedFamilies []SelectedFamily
}

type RenderResult struct {
	Message  string
	Evidence []string
	Notes    []string
}

type ExperimentalConfig struct {
	Provider    string
	Model       string
	Profile     string
	Prompt      string
	AdapterMode ExperimentalAdapterMode
	Enabled     bool
}

type Config struct {
	DevType      DevType
	Jargon       JargonLevel
	Complexity   Complexity
	Duration     int
	Alerts       bool
	Project      string
	Minimal      bool
	Team         bool
	Framework    string
	Seed         int64
	OutputFormat OutputFormat
	NoColor      bool
	Trace        bool
	Experimental ExperimentalConfig
}

type rawConfig struct {
	DevType      string
	Jargon       string
	Complexity   string
	Duration     int
	Alerts       bool
	Project      string
	Minimal      bool
	Team         bool
	Framework    string
	Seed         int64
	OutputFormat string
	NoColor      bool
	Trace        bool
	Experimental ExperimentalConfig
}

type ValueCatalog struct {
	Config struct {
		DevType      []string `json:"devType"`
		Jargon       []string `json:"jargon"`
		Complexity   []string `json:"complexity"`
		OutputFormat []string `json:"outputFormat"`
	} `json:"config"`
	Flags struct {
		Alerts  []bool `json:"alerts"`
		Team    []bool `json:"team"`
		Minimal []bool `json:"minimal"`
		NoColor []bool `json:"noColor"`
		Trace   []bool `json:"trace"`
	} `json:"flags"`
	Experimental struct {
		Provider    []string `json:"provider"`
		AdapterMode []string `json:"adapterMode"`
	} `json:"experimental"`
	Families []FamilySpec `json:"families"`
}

func (f Family) String() string { return string(f) }

func (g FamilyGroup) String() string { return string(g) }

func (d DevType) String() string { return string(d) }

func (j JargonLevel) String() string { return string(j) }

func (c Complexity) String() string { return string(c) }

func (o OutputFormat) String() string { return string(o) }

func (m ExperimentalAdapterMode) String() string { return string(m) }

func formatUnsupportedExperimentalMessage(cfg ExperimentalConfig) string {
	parts := []string{"experimental provider mode is not implemented in go-stakeholder"}
	if cfg.Provider != "" {
		parts = append(parts, fmt.Sprintf("provider=%s", cfg.Provider))
	}
	if cfg.Model != "" {
		parts = append(parts, fmt.Sprintf("model=%s", cfg.Model))
	}
	if cfg.Profile != "" {
		parts = append(parts, fmt.Sprintf("profile=%s", cfg.Profile))
	}
	if cfg.Prompt != "" {
		parts = append(parts, fmt.Sprintf("prompt=%s", cfg.Prompt))
	}
	if cfg.AdapterMode != "" {
		parts = append(parts, fmt.Sprintf("adapter-mode=%s", cfg.AdapterMode))
	}
	return fmt.Sprintf("%s; flags are parsed but remain fail-fast", joinParts(parts))
}

func joinParts(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	out := parts[0]
	for _, part := range parts[1:] {
		out += "; " + part
	}
	return out
}

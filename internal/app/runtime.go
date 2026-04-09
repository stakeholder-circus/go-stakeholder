package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	cfg, listValues, err := parseArgs(args, stderr)
	if err != nil {
		return 2
	}
	if cfg.Experimental.Enabled {
		_, _ = fmt.Fprintln(stderr, formatUnsupportedExperimentalMessage(cfg.Experimental))
		return 2
	}
	if listValues {
		return writeValueCatalog(stdout)
	}
	return writeSession(cfg, stdout, stderr)
}

func parseArgs(args []string, stderr io.Writer) (Config, bool, error) {
	fs := flag.NewFlagSet("stakeholder", flag.ContinueOnError)
	fs.SetOutput(stderr)

	listValues := fs.Bool("list-values", false, "")
	devType := fs.String("dev-type", DevTypeFullStack.String(), "")
	jargon := fs.String("jargon", JargonLevelMedium.String(), "")
	complexity := fs.String("complexity", ComplexityMedium.String(), "")
	duration := fs.Int("duration", 0, "")
	alerts := fs.Bool("alerts", false, "")
	project := fs.String("project", "", "")
	minimal := fs.Bool("minimal", false, "")
	team := fs.Bool("team", false, "")
	framework := fs.String("framework", "", "")
	seed := fs.String("seed", "0", "")
	outputFormat := fs.String("output-format", OutputFormatText.String(), "")
	noColor := fs.Bool("no-color", false, "")
	trace := fs.Bool("trace", false, "")
	experimentalProvider := fs.String("experimental-provider", "", "")
	experimentalModel := fs.String("experimental-model", "", "")
	experimentalProfile := fs.String("experimental-profile", "", "")
	experimentalPrompt := fs.String("experimental-prompt", "", "")
	experimentalAdapterMode := fs.String("experimental-adapter-mode", "", "")

	if err := fs.Parse(args); err != nil {
		return Config{}, false, err
	}

	seedValue, err := strconv.ParseInt(strings.TrimSpace(*seed), 10, 64)
	if err != nil {
		return Config{}, false, fmt.Errorf("invalid seed %q: %w", *seed, err)
	}

	raw := rawConfig{
		DevType:      *devType,
		Jargon:       *jargon,
		Complexity:   *complexity,
		Duration:     *duration,
		Alerts:       *alerts,
		Project:      *project,
		Minimal:      *minimal,
		Team:         *team,
		Framework:    *framework,
		Seed:         seedValue,
		OutputFormat: *outputFormat,
		NoColor:      *noColor,
		Trace:        *trace,
		Experimental: ExperimentalConfig{
			Provider:    *experimentalProvider,
			Model:       *experimentalModel,
			Profile:     *experimentalProfile,
			Prompt:      *experimentalPrompt,
			AdapterMode: ExperimentalAdapterMode(*experimentalAdapterMode),
		},
	}

	cfg, err := normalizeConfig(raw)
	if err != nil {
		return Config{}, false, err
	}

	return cfg, *listValues, nil
}

func normalizeConfig(raw rawConfig) (Config, error) {
	devType, err := normalizeDevType(raw.DevType)
	if err != nil {
		return Config{}, err
	}
	jargon, err := normalizeJargon(raw.Jargon)
	if err != nil {
		return Config{}, err
	}
	complexity, err := normalizeComplexity(raw.Complexity)
	if err != nil {
		return Config{}, err
	}
	outputFormat, err := normalizeOutputFormat(raw.OutputFormat)
	if err != nil {
		return Config{}, err
	}
	experimental := normalizedExperimentalConfig(raw.Experimental)
	if err := validateExperimentalConfig(experimental); err != nil {
		return Config{}, err
	}
	return Config{
		DevType:      devType,
		Jargon:       jargon,
		Complexity:   complexity,
		Duration:     raw.Duration,
		Alerts:       raw.Alerts,
		Project:      strings.TrimSpace(raw.Project),
		Minimal:      raw.Minimal,
		Team:         raw.Team,
		Framework:    strings.TrimSpace(raw.Framework),
		Seed:         raw.Seed,
		OutputFormat: outputFormat,
		NoColor:      raw.NoColor,
		Trace:        raw.Trace,
		Experimental: experimental,
	}, nil
}

func normalizeDevType(value string) (DevType, error) {
	if value == "" {
		return DevTypeFullStack, nil
	}
	switch DevType(strings.ToLower(strings.TrimSpace(value))) {
	case DevTypeBackend, DevTypeBlockchain, DevTypeDataScience, DevTypeDevOps, DevTypeFrontend, DevTypeFullStack, DevTypeGameDevelopment, DevTypeMachineLearning, DevTypeSecurity, DevTypeSystemsProgramming:
		return DevType(strings.ToLower(strings.TrimSpace(value))), nil
	default:
		return "", fmt.Errorf("invalid dev-type %q", value)
	}
}

func normalizeJargon(value string) (JargonLevel, error) {
	if value == "" {
		return JargonLevelMedium, nil
	}
	switch JargonLevel(strings.ToLower(strings.TrimSpace(value))) {
	case JargonLevelLow, JargonLevelMedium, JargonLevelHigh, JargonLevelExtreme:
		return JargonLevel(strings.ToLower(strings.TrimSpace(value))), nil
	default:
		return "", fmt.Errorf("invalid jargon %q", value)
	}
}

func normalizeComplexity(value string) (Complexity, error) {
	if value == "" {
		return ComplexityMedium, nil
	}
	switch Complexity(strings.ToLower(strings.TrimSpace(value))) {
	case ComplexityLow, ComplexityMedium, ComplexityHigh, ComplexityExtreme:
		return Complexity(strings.ToLower(strings.TrimSpace(value))), nil
	default:
		return "", fmt.Errorf("invalid complexity %q", value)
	}
}

func normalizeOutputFormat(value string) (OutputFormat, error) {
	if value == "" {
		return OutputFormatText, nil
	}
	switch OutputFormat(strings.ToLower(strings.TrimSpace(value))) {
	case OutputFormatText, OutputFormatJSON:
		return OutputFormat(strings.ToLower(strings.TrimSpace(value))), nil
	default:
		return "", fmt.Errorf("invalid output-format %q", value)
	}
}

func writeValueCatalog(stdout io.Writer) int {
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(valueCatalog()); err != nil {
		return 1
	}
	return 0
}

func writeSession(cfg Config, stdout io.Writer, stderr io.Writer) int {
	events := BuildSession(cfg)
	switch cfg.OutputFormat {
	case OutputFormatJSON:
		encoder := json.NewEncoder(stdout)
		encoder.SetEscapeHTML(false)
		for _, event := range events {
			if err := encoder.Encode(event); err != nil {
				_, _ = fmt.Fprintln(stderr, err)
				return 1
			}
		}
		return 0
	default:
		for _, line := range formatTextStream(events, cfg.Minimal, cfg.NoColor) {
			if _, err := fmt.Fprintln(stdout, line); err != nil {
				_, _ = fmt.Fprintln(stderr, err)
				return 1
			}
		}
		return 0
	}
}

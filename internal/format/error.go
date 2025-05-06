package format

import (
	"errors"
	"strings"

	je "github.com/ymatsukawa/jak/internal/errors"
)

// FormatError formats an error for display with context and help message
func FormatError(err error) string {
	if err == nil {
		return ""
	}

	var buffer strings.Builder

	// Divider line
	divider := strings.Repeat("─", 50)
	buffer.WriteString(ColorizeError(divider) + "\n")

	// Error title and message
	buffer.WriteString(ColorizeError("ERROR: ") + err.Error() + "\n\n")

	// Add helpful context based on error type
	var helpMsg string
	if errors.Is(err, je.ErrInvalidURL) {
		helpMsg = "The URL format is invalid. Make sure it includes scheme (http:// or https://) and host."
	} else if errors.Is(err, je.ErrInvalidMethod) {
		helpMsg = "Invalid HTTP method. Supported methods are: GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS."
	} else if errors.Is(err, je.ErrInvalidConfig) {
		helpMsg = "The configuration file is not valid or could not be loaded. Check the file format and permissions."
	} else if errors.Is(err, je.ErrConfigValidation) {
		helpMsg = "The configuration failed validation. Check that all required fields are present and correct."
	} else if errors.Is(err, je.ErrRequestExecution) {
		helpMsg = "Failed to execute the HTTP request. Check your network connection and the server status."
	} else if errors.Is(err, je.ErrRequestCreation) {
		helpMsg = "Failed to create the HTTP request. Check your request parameters."
	} else if errors.Is(err, je.ErrInvalidHeader) {
		helpMsg = "Invalid header format. Headers must be in the format 'Key: Value'."
	} else if errors.Is(err, je.ErrInvalidBody) {
		helpMsg = "Invalid request body. Check the format of your JSON, form, or raw body."
	} else if errors.Is(err, je.ErrResponseRead) {
		helpMsg = "Failed to read the response. The server may have returned an invalid response."
	}

	if helpMsg != "" {
		buffer.WriteString(ColorizeInfo("HELP: ") + helpMsg + "\n")
	}

	// Add command hint for some common errors
	if isCommandRelatedError(err) {
		buffer.WriteString("\n" + ColorizeHeader("TIP: ") + "Run 'jak --help' for usage information.\n")
	}

	// Divider line
	buffer.WriteString(ColorizeError(divider) + "\n")

	return buffer.String()
}

// FormatCommandError formats command-line usage errors
func FormatCommandError(cmdName string, err error) string {
	if err == nil {
		return ""
	}

	var buffer strings.Builder

	// Divider line
	divider := strings.Repeat("─", 50)
	buffer.WriteString(ColorizeError(divider) + "\n")

	// Error message
	buffer.WriteString(ColorizeError("COMMAND ERROR: ") + err.Error() + "\n\n")

	// Add usage help
	buffer.WriteString(ColorizeHeader("USAGE: ") + "\n")

	switch cmdName {
	case "req":
		buffer.WriteString("  jak req [method] [url] [flags]\n\n")
		buffer.WriteString("Examples:\n")
		buffer.WriteString("  jak req GET https://example.com\n")
		buffer.WriteString("  jak req POST https://api.example.com/data -H \"Content-Type: application/json\" -j '{\"key\":\"value\"}'\n")
	case "bat":
		buffer.WriteString("  jak bat [config_file]\n\n")
		buffer.WriteString("Examples:\n")
		buffer.WriteString("  jak bat config.toml\n")
	case "chain":
		buffer.WriteString("  jak chain [config_file]\n\n")
		buffer.WriteString("Examples:\n")
		buffer.WriteString("  jak chain config.toml\n")
	default:
		buffer.WriteString("  jak [command] [args] [flags]\n\n")
		buffer.WriteString("Available Commands:\n")
		buffer.WriteString("  req     Execute a simple HTTP request\n")
		buffer.WriteString("  bat     Execute batch requests from a config file\n")
		buffer.WriteString("  chain   Execute chain requests with dependencies\n")
	}

	buffer.WriteString("\nRun 'jak --help' or 'jak [command] --help' for more information.\n")

	// Divider line
	buffer.WriteString(ColorizeError(divider) + "\n")

	return buffer.String()
}

// isCommandRelatedError checks if the error is related to command usage
func isCommandRelatedError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return strings.Contains(errMsg, "required argument") ||
		strings.Contains(errMsg, "flag") ||
		strings.Contains(errMsg, "argument") ||
		strings.Contains(errMsg, "unknown command") ||
		strings.Contains(errMsg, "invalid syntax")
}

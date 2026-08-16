package terminal

// DetectBackground exposes the file-driven detection core so tests can exercise
// the nil, non-terminal, and pipe paths without a real terminal.
var DetectBackground = detectBackground

// DetectTrueColor exposes the env-driven detection core so tests can exercise
// each environment without mutating the process or hitting the cached result.
var DetectTrueColor = detectTrueColor

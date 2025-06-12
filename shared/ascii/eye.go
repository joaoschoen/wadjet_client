package ascii

import "strings"

func Eye() string {
	eye := "y               ███████████████                       \n"
	eye += "y          █████████████████████████                  \n"
	eye += "y       ██████                   █████                \n"
	eye += "y    ██████      ████████████       █████             \n"
	eye += "y  █████      █████       ███████      █████          \n"
	eye += "y█████     ██████  eb███████ey   ███████      ██████      \n"
	eye += "y███     █████    eb█████████ey      █████       █████████\n"
	eye += "y      █████     eb███████████ey        █████             \n"
	eye += "y     ███        eb███████████ey           ███████████    \n"
	eye += "y    ████████     eb█████████ey        ███████████████████\n"
	eye += "y    ██   ████████  eb█████ey  █████████████              \n"
	eye += "y              ██████████████████                     \n"
	eye += "y               ██████                                \n"
	eye += "y           ███████████                               \n"
	eye += "y           ████████████         eb┓ ┏┏┓┳┓┏┳┏┓┏┳┓ey       \n"
	eye += "y               ████ ████        eb┃┃┃┣┫┃┃ ┃┣  ┃ ey       \n"
	eye += "y                ███  ████       eb┗┻┛┛┗┻┛┗┛┗┛ ┻ ey       \n"
	eye += "y                ███   ████                           \n"
	eye += "y                ███     ████                         \n"
	eye += "y                ███      ████                        \n"
	eye += "y                 ██        ███                       \n"
	eye += "y                 ██         ████                ████ \n"
	eye += "y                 ██           ████             ██████\n"
	eye += "y                ████            ████            █  ██\n"
	eye += "y                                 █████            ███\n"
	eye += "y                                    ███████     ████ \n"
	eye += "y                                        ██████████   e\n"

	// COLOR CODES:
	var text_color = map[string]string{
		"yellow": "\033[33m",
		"blue":   "\033[34m",
	}
	const end = "\033[0m"
	eye = strings.ReplaceAll(eye, "y", text_color["yellow"])
	eye = strings.ReplaceAll(eye, "b", text_color["blue"])
	eye = strings.ReplaceAll(eye, "e", end)
	return eye
}

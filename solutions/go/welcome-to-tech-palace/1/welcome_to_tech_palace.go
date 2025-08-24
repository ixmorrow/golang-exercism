package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return("Welcome to the Tech Palace, " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var sb strings.Builder // Zero value is ready to use
    for i:= 0; i < numStarsPerLine; i++ {
        sb.WriteByte('*')
    }
    sb.WriteByte('\n')

    sb.WriteString(welcomeMsg)
    sb.WriteByte('\n')
    
    for i:= 0; i < numStarsPerLine; i++ {
        sb.WriteByte('*')
    }

	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var sb strings.Builder
    for _, r := range oldMsg {
        if r != '*' {
            sb.WriteRune(r)
        }
    }
    return(strings.TrimSpace(sb.String()))
}

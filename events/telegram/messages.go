package telegram

const msgHelp = `I can save and keep you cinema links. Also I can offer you them to see.

In order to save the cinema, just send me al link to it.

In order to get a random cinema from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!`

const msgHello = "Hi there! π\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command π€"
	msgNoSavedPages   = "You have no saved cinemas π"
	msgSaved          = "Saved! π"
	msgAlreadyExists  = "You have already have this cinema in your list πΆβπ«"
)

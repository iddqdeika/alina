# alina

Important! Package is currently at work and hasn't any tests yet.

Effort to implement usable VK api.
Alina uses longpollAPI to get new events. 
Now we can only use message handler.

Use alina.New to create new Alina.

Alina can take message handlers to handle incoming messages.

Alina can give simple MessagesApi (alina.GetMEssagesApi), that implements methods to send simple messages and get messagehistory for given conversation.

We already can, for example, add MessageHandler to answer for some messages:

alina.AddMessageHandler(func(message definitions.PrivateMessage, e error) {

		if err != nil {
			logger.Error(err)
			return
		}
		if strings.Contains(message.GetText(), "the best wife") {
			alina.GetMessagesApi().SendSimpleMessage(strconv.Itoa(message.GetPeerId()), "Alina, of course!")
		}
})


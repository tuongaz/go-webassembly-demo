package engine

type Platform interface {
	// GameReady Init is called to notify the platform the game is ready
	GameReady(*Board)

	// Input returns a channel of events sent from the platform
	Input() chan Input
	
	// Output Render is called to notify the platform to render the board
	Output(*Board)
}

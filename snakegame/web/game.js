let app
let gameScreen

const stepSize = 30
const readyTextColor = "#FC5185"
const readyTextColorShadow = "#8B2D4F"
const boxColorAlternative = "0x7aadff"
const boxColor = "0x86b3fc"

// is called from Go
function gameReady(boardJSON) {
    const board = JSON.parse(boardJSON)
    if (!app) {
        initApp(board)
    }

    drawReadyScreen(board)
}

input()
// calling Go func
// changeDirection func is defined in Go
function input() {
    window.addEventListener('keydown', e => {
        if (e.key === 'ArrowRight' || e.key === 'd') {
            changeDirection("right");
        } else if (e.key === 'ArrowLeft' || e.key === 'a') {
            changeDirection("left");
        } else if (e.key === 'ArrowUp' || e.key === 'w') {
            changeDirection("up");
        } else if (e.key === 'ArrowDown' || e.key === 's') {
            changeDirection("down");
        }
    });
}

// is called from Go
function output(data) {
    const b = JSON.parse(data)
    drawBoard(b)
}

function drawReadyScreen(board) {
    gameScreen.removeChildren();

    drawComponents(board);

    const readyText = new PIXI.Text("Ready!", new PIXI.TextStyle({
        fill: [readyTextColor],
        fontSize: 40,
        dropShadow: true,
        dropShadowColor: readyTextColorShadow,
        dropShadowBlur: 2,
        dropShadowAngle: Math.PI / 6,
        dropShadowDistance: 2,
    }))
    readyText.x = 160
    readyText.y = 60
    gameScreen.addChild(readyText);
}


function drawBoard(board) {
    gameScreen.removeChildren();
    drawComponents(board);
}

function drawComponents(board) {
    const snake = board.snake;
    drawSnake(snake)
    drawFood(board.foods)
    drawPoints(snake.points)
}

function drawSnake(snake) {
    let snakeHead = PIXI.Sprite.from("head.png")
    snakeHead.anchor.set(0.5);
    updateSnakeHeadDirection(snakeHead, snake.direction)

    for (let i = 0; i < snake.body.length; i++) {
        let part = snake.body[i];

        if (i === snake.body.length - 1) { // This is the last part - the head
            drawSnakeHead(snakeHead, part)
        } else {
            drawSnakeBody(part)
        }
    }
}

function updateSnakeHeadDirection(snakeHead, direction) {
    switch (direction) {
        case 'up':
            snakeHead.rotation = 0;
            break;
        case 'down':
            snakeHead.rotation = Math.PI;  // 180 degrees
            break;
        case 'left':
            snakeHead.rotation = -Math.PI / 2;  // -90 degrees
            break;
        case 'right':
            snakeHead.rotation = Math.PI / 2;  // 90 degrees
            break;
    }
}

function drawPoints(points) {
    const pointText = new PIXI.Text("Points: " + points, new PIXI.TextStyle({
        fill: ['#ffffff'],
        fontSize: 25,
        dropShadow: true,
        dropShadowColor: '#888888',
        dropShadowBlur: 2,
        dropShadowAngle: Math.PI / 6,
        dropShadowDistance: 2,
    }))
    pointText.x = 10
    pointText.y = 10
    gameScreen.addChild(pointText);
}

function drawSnakeBody(part) {
    let block = new PIXI.Graphics();
    block.beginFill(0x00a341);
    block.drawCircle(part.x * stepSize + stepSize / 2, part.y * stepSize + stepSize / 2, stepSize / 2 - 1);
    block.endFill();
    gameScreen.addChild(block);
}

function drawSnakeHead(snakeHead, part) {
    const snakeHeadPadding = 5
    snakeHead.x = part.x * stepSize + stepSize / 2;
    snakeHead.y = part.y * stepSize + stepSize / 2;
    snakeHead.width = stepSize + snakeHeadPadding;
    snakeHead.height = stepSize + snakeHeadPadding;
    gameScreen.addChild(snakeHead);
}

function drawFood(foods) {
    for (let food of foods) {
        const emoji = new PIXI.Text(String.fromCodePoint(food.image), new PIXI.TextStyle({
                fontSize: stepSize - 3,
            })
        )
        emoji.x = food.coord.x * stepSize
        emoji.y = food.coord.y * stepSize
        gameScreen.addChild(emoji);
    }
}

function initApp(board) {
    app = new PIXI.Application({
        width: board.width * stepSize,
        height: board.height * stepSize,
        antialias: true,
        transparent: false,
        resolution: 1
    })
    document.body.appendChild(app.view);
    drawGameBackground(board)

    gameScreen = new PIXI.Container();
    app.stage.addChild(gameScreen);
}

function drawGameBackground(board) {
    let backgroundContainer = new PIXI.Container();
    app.stage.addChild(backgroundContainer);

    for (let i = 0; i < board.width; i++) {
        for (let j = 0; j < board.height; j++) {
            const square = new PIXI.Graphics();
            square.beginFill((i + j) % 2 === 0 ? boxColor : boxColorAlternative);
            square.drawRect(j * stepSize, i * stepSize, stepSize, stepSize);
            square.endFill();
            backgroundContainer.addChild(square);
        }
    }
}

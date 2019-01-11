const gameArea = {
	canvas: document.createElement('canvas'),
	start: function(width, height) {
		this.canvas.width = width;
		this.canvas.height = height;
		// this.canvas.style.border='2px solid black';
		this.ctx = this.canvas.getContext('2d');
		document.body.insertBefore(this.canvas, document.body.childNodes[1]);
	}
};

let id = document.createElement('p');
id.style.textAlign = 'center';
id.style.font = '3em Arial';
id.style.color = '#008cba';
id.innerHTML = 'Game ID: ' + document.location.pathname.replace('/id/', '');
document.body.insertBefore(id, document.body.childNodes[0]);

gameArea.start(500, 500);
initCanvas();

function initCanvas() {
	gameArea.ctx.beginPath();
	gameArea.ctx.rect(gameArea.canvas.width/3-5, 20, 10, gameArea.canvas.height-40);
	gameArea.ctx.fillStyle = 'black';
	gameArea.ctx.fill()

	gameArea.ctx.beginPath();
	gameArea.ctx.rect(2*gameArea.canvas.width/3-5, 20, 10, gameArea.canvas.height-40);
	gameArea.ctx.fillStyle = 'black';
	gameArea.ctx.fill()

	gameArea.ctx.beginPath();
	gameArea.ctx.rect(20, gameArea.canvas.height/3-5, gameArea.canvas.height-40, 10);
	gameArea.ctx.fillStyle = 'black';
	gameArea.ctx.fill()

	gameArea.ctx.beginPath();
	gameArea.ctx.rect(20, 2*gameArea.canvas.height/3-5, gameArea.canvas.height-40, 10);
	gameArea.ctx.fillStyle = 'black';
	gameArea.ctx.fill()
}


function drawX(x, y) {
	gameArea.ctx.beginPath();
	gameArea.ctx.font = 2*gameArea.canvas.width/7 +'px Arial';
	gameArea.ctx.fillStyle = '#2874A6';
	gameArea.ctx.fillText('X', gameArea.canvas.width/6 - 1.7*gameArea.canvas.width/18 + x*gameArea.canvas.width/3, gameArea.canvas.height/6 + 2*gameArea.canvas.width/18 + y*gameArea.canvas.height/3);
}

function drawO(x, y) {
	gameArea.ctx.beginPath();
	gameArea.ctx.font = 2*gameArea.canvas.width/7 +'px Arial';

	gameArea.ctx.fillStyle = "#7B241C"; 
	gameArea.ctx.fillText('O', gameArea.canvas.width/6 - 2*gameArea.canvas.width/18 + x*gameArea.canvas.width/3, gameArea.canvas.height/6 + 2*gameArea.canvas.width/18 + y*gameArea.canvas.height/3);
}

addEventListener('click', handleClick, false);

// Get player id
// TODO: Use cookies
var playernum;
window.onload = function() {
	const params = JSON.stringify({id: window.location.pathname.replace('/id/', '')});
	console.log(params);
	getJSON('/getid/', params,
		function (err, data) {
			console.log(data);
			if (err !== null) {
				console.log('Error retrieving data: ' + err);
			} else {
				playernum = data.playernum;
			}
		});
}

function handleClick(e) {
	let x = e.clientX;
	let y = e.clientY - e.target.getBoundingClientRect().top;
	let X, Y;
	if (x < gameArea.canvas.width/3) {
		X = 0;
	} else if (x < 2*gameArea.canvas.width/3) {
		X = 1;
	} else if (x < gameArea.canvas.width) {
		X = 2;
	}

	if (y < gameArea.canvas.height/3) {
		Y = 0;
	} else if (y < 2*gameArea.canvas.height/3) {
		Y = 1;
	} else if (y < gameArea.canvas.height) {
		Y = 2;
	}
	const params = JSON.stringify({player: playernum, move: [X,Y]});
	sendMove(window.location.pathname, params, {X: X, Y: Y});
}

let sendMove = function (url, params, move) {
	getJSON(url, params,
		function (err, data) {
			if (err !== null) {
				console.log('Error retrieving data: ' + err);
			} else {
				if (data.valid === true) {
					updateCanvas(data);
				}
			}
		});
};

let updateCanvas = function (data) {
	/*please run my child*/gameArea.ctx.clearRect(0, 0, gameArea.canvas.width, gameArea.canvas.Height);
	initCanvas();
	for (i = 0; i < data.board.length; i++) {
		for (j = 0; j < data.board[0].length; j++) {
			if (data.board[i][j] == 1) {
				drawX(i, j);
			} else if (data.board[i][j] == 2) {
				drawO(i, j);
			}
		}
	}
	if (data.winner !== 0) {
		console.log('Winner is player ' + data.winner + '!');
		clearInterval(gameLoopID);
		return;
	}
};

var gameLoopID = setInterval(function() {
	getJSON(window.location.pathname, null,
		function (err, data) {
			if (err !== null) {
				console.log('Error retrieving data: ' + err);
			} else {
				updateCanvas(data);
			}
		});
}, 250);

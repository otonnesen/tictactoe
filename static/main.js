const gameArea = {
	canvas: document.createElement('canvas'),
	start: function(width, height) {
		this.canvas.width = width;
		this.canvas.height = height;
		// this.canvas.style.border='2px solid black';
		this.ctx = this.canvas.getContext('2d');
		document.body.insertBefore(this.canvas, document.body.childNodes[0]);
	},
	components: []
};

gameArea.start(750, 750);
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

function drawX(x, y) {
	gameArea.ctx.beginPath();
	gameArea.ctx.font = 2*gameArea.canvas.width/7 +'px Arial';
	gameArea.ctx.fillText('X', gameArea.canvas.width/6 - 1.7*gameArea.canvas.width/18 + x*gameArea.canvas.width/3, gameArea.canvas.height/6 + 2*gameArea.canvas.width/18 + y*gameArea.canvas.height/3);
}

function drawO(x, y) {
	gameArea.ctx.beginPath();
	gameArea.ctx.font = 2*gameArea.canvas.width/7 +'px Arial';
	gameArea.ctx.fillText('O', gameArea.canvas.width/6 - 2*gameArea.canvas.width/18 + x*gameArea.canvas.width/3, gameArea.canvas.height/6 + 2*gameArea.canvas.width/18 + y*gameArea.canvas.height/3);
}

function getMousePos(canvas, e) {
	var rect = canvas.getBoundingClientRect();
	return {
		x: e.clientX - rect.left,
		y: e.clientY - rect.top
	};
}

drawO(0, 2
drawX(1, 2);

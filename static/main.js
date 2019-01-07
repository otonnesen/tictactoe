const gameArea = {
	canvas: document.createElement('canvas'),
	start: function(width, height) {
		this.canvas.width = width;
		this.canvas.height = height;
		this.canvas.style.border='2px solid black';
		this.ctx = this.canvas.getContext('2d');
		document.body.insertBefore(this.canvas, document.body.childNodes[0]);
	},
	components: []
};

gameArea.start(500, 500);
gameArea.ctx.beginPath();
gameArea.ctx.rect(50, 50, 10, 10);
gameArea.ctx.fillStyle = 'black';
gameArea.ctx.fill()

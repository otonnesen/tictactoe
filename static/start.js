const startArea = {
	canvas: document.createElement('canvas'),
	start: function(width, height) {
		this.canvas.width = width;
		this.canvas.height = height;
		// this.canvas.style.border='2px solid black';
        this.ctx = this.canvas.getContext('2d');
		document.body.insertBefore(this.canvas, document.body.childNodes[0]);
	}
};

startArea.start(1200,900);

var ctx = startArea.canvas.getContext('2d');

function roundRect(ctx, x, y, width, height, radius, fill, stroke) {
    if (typeof stroke == "undefined" ) {
      stroke = true;
    }
    if (typeof radius === "undefined") {
      radius = 5;
    }
    ctx.beginPath();
    ctx.moveTo(x + radius, y);
    ctx.lineTo(x + width - radius, y);
    ctx.quadraticCurveTo(x + width, y, x + width, y + radius);
    ctx.lineTo(x + width, y + height - radius);
    ctx.quadraticCurveTo(x + width, y + height, x + width - radius, y + height);
    ctx.lineTo(x + radius, y + height);
    ctx.quadraticCurveTo(x, y + height, x, y + height - radius);
    ctx.lineTo(x, y + radius);
    ctx.quadraticCurveTo(x, y, x + radius, y);
    ctx.closePath();
    if (stroke) {
      ctx.stroke();
    }
    if (fill) {
      ctx.fill();
    }        
  }
ctx.lineWidth = 4;
ctx.strokeStyle = "#000000";
ctx.fillStyle = "#abc";
roundRect(ctx, 10, 10, 150, 50, 10, true);
ctx.font="20px Georgia";
ctx.textAlign="center"; 
ctx.textBaseline = "middle";
ctx.fillStyle = "#000000";
var rectHeight = 50;
var rectWidth = 150;
var rectX = 10;
var rectY = 10;
ctx.fillText("Click to Start",rectX+(rectWidth/2),rectY+(rectHeight/2));


  

document.addEventListener('click', handleClick, false);


function handleClick(e) {
    console.log('('+e.clientX+','+e.clientY+')');
    let x = e.clientX;
    let y = e.clientY;
    if (x > 10 && x < 170) {
		if(y > 10 && y < 70){
            //start button pressed
            console.log('START');
        }
	}
}



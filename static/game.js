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
	gameArea.ctx.fillStyle = "#2874A6";
	gameArea.ctx.fillText('X', gameArea.canvas.width/6 - 1.7*gameArea.canvas.width/18 + x*gameArea.canvas.width/3, gameArea.canvas.height/6 + 2*gameArea.canvas.width/18 + y*gameArea.canvas.height/3);
}

function drawO(x, y) {
	gameArea.ctx.beginPath();
	gameArea.ctx.font = 2*gameArea.canvas.width/7 +'px Arial';
	gameArea.ctx.fillStyle = "#7B241C"; 
	gameArea.ctx.fillText('O', gameArea.canvas.width/6 - 2*gameArea.canvas.width/18 + x*gameArea.canvas.width/3, gameArea.canvas.height/6 + 2*gameArea.canvas.width/18 + y*gameArea.canvas.height/3);
}

document.addEventListener('click', handleClick, false);

function handleClick(e) {
	console.log('('+e.clientX+','+e.clientY+')');
	let x = e.clientX;
	let y = e.clientY;
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
	drawO(X, Y);
	return {
		x: X,
		y: Y,
	}
	
}

//taken from loanshark

let getJSON = function (url, params, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-type", "application/json; charset=utf-8");

    xhr.responseType = 'json';
    xhr.onload = function () {
        const status = xhr.status;
        if (status === 200) {
            callback(null, xhr.response);
        } else {
            callback(status, xhr.response);
        }
    };
    xhr.send(params);
};

let getData = function (url, params) {
    getJSON(url, params,
        function (err, data) {
            if (err !== null) {
                console.log("Error retrieving data: " + err);
            } else {
                login(data['validate']);
            }
            //document.getElementById('submit').style.cursor = '';
        });
};

(function() {

  var VIEW_SIZE = 120;
  var LINE_WIDTH = 12;

  function PathView(path) {
    this._ratio = Math.ceil(window.devicePixelRatio);
    this._element = document.createElement('canvas');
    this._element.width = this._ratio * VIEW_SIZE;
    this._element.height = this._ratio * VIEW_SIZE;
    this._element.className = 'line-drawing';
    if (path) {
      this._path = path;
      this._draw();
    } else {
      this._path = [];
      this._registerMouseEvents();
    }
    this.onDrawn = null;
  }

  PathView.prototype.element = function() {
    return this._element;
  };

  PathView.prototype.path = function() {
    return this._path;
  };

  PathView.prototype._draw = function() {
    if (this._path.length === 0) {
      return;
    }

    var ctx = this._element.getContext('2d');
    ctx.clearRect(0, 0, this._element.width, this._element.height);

    ctx.save();
    ctx.scale(this._ratio, this._ratio);
    ctx.fillStyle = 'black';
    ctx.strokeStyle = 'black';
    ctx.lineJoin = 'round';
    ctx.lineCap = 'round';
    ctx.lineWidth = LINE_WIDTH;

    if (this._path.length === 1) {
      ctx.beginPath();
      ctx.arc(this._path[0].x, this._path[0].y, ctx.lineWidth/2, 0, Math.PI*2);
      ctx.closePath();
      ctx.fill();
    } else {
      ctx.beginPath();
      ctx.moveTo(this._path[0].x, this._path[0].y);
      for (var i = 1, len = this._path.length; i < len; ++i) {
        ctx.lineTo(this._path[i].x, this._path[i].y);
      }
      ctx.stroke();
    }

    ctx.restore();
  };

  PathView.prototype._registerMouseEvents = function() {
    var downListener = function(e) {
      var upListener, moveListener;
      upListener = function() {
        window.removeEventListener('mouseup', upListener);
        window.removeEventListener('mouseleave', upListener);
        window.removeEventListener('mousemove', moveListener);
        window.removeEventListener('mousedown', downListener);
        this.onDrawn();
      }.bind(this);

      moveListener = function(e) {
        this._path.push(this._mousePosition(e));
        this._draw();
      }.bind(this);

      this._path.push(this._mousePosition(e));
      this._draw();
      window.addEventListener('mouseup', upListener);
      window.addEventListener('mouseleave', upListener);
      window.addEventListener('mousemove', moveListener);
    }.bind(this);

    this._element.addEventListener('mousedown', downListener);
  };

  PathView.prototype._mousePosition = function(e) {
    var x = e.pageX - this._element.offsetLeft;
    var y = e.pageY - this._element.offsetTop;
    return {x: x, y: y};
  };

  window.app.PathView = PathView;

})();

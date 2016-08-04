(function() {

  window.app = {};

  var pathViews = [];

  function initialize() {
    setupNextCanvas();
    var expBut = document.getElementById('export-button');
    expBut.onclick = function() {
      var paths = [];
      for (var i = 0, len = pathViews.length; i < len; ++i) {
        paths.push(pathViews[i].path());
      }
      document.body.textContent = '';
      var pre = document.createElement('pre');
      pre.textContent = JSON.stringify(paths);
      pre.style.textAlign = 'left';
      pre.style.whiteSpace = 'pre-wrap';
      document.body.appendChild(pre);
    };
  }

  function setupNextCanvas() {
    var container = document.getElementById('draw-container');
    var view = new window.app.PathView();
    container.appendChild(view.element());
    view.onDrawn = function() {
      container.removeChild(view.element());
      addedDrawing(view);
      setupNextCanvas();
    };
  }

  function addedDrawing(pathView) {
    var listing = document.getElementById('listing');

    var row = document.createElement('tr');
    var drawingCol = document.createElement('td');
    var deleteCol = document.createElement('td');

    drawingCol.appendChild(pathView.element());

    var deleteButton = document.createElement('button');
    deleteButton.textContent = 'Delete';
    deleteButton.onclick = function() {
      listing.removeChild(row);
      pathViews.splice(pathViews.indexOf(pathView), 1);
    };
    deleteCol.appendChild(deleteButton);

    row.appendChild(drawingCol);
    row.appendChild(deleteCol);
    listing.appendChild(row);

    pathViews.push(pathView);
  }

  window.addEventListener('load', initialize);

})();

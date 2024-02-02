var canvas = new fabric.Canvas('canvas');


// Path to your SVG file
var svgFilePath = 'artworks/user-drawing.svg';

// Fetch the SVG file using the fetch API
fetch(svgFilePath)
    .then(response => response.text())
    .then(svgString => {
        // Load SVG string onto the canvas
        fabric.loadSVGFromString(svgString, function (objects, options) {
            var svgObject = fabric.util.groupSVGElements(objects, options);
            canvas.add(svgObject);
            canvas.selection = false; // make canvas not selectable/editable
            canvas.renderAll();
        });
    }).catch(error => console.error('Error fetching SVG file:', error));


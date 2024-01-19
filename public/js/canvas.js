const canvas = new fabric.Canvas('canvas', {
    width: 300,
    height: 450,
    backgroundColor: 'white'

});

/*  SAMPLE SHAPES
var circle = new fabric.Circle({
    radius: 20, fill: 'green', left: 100, top: 100
});

var triangle = new fabric.Triangle({
    width: 20, height: 30, fill: 'blue', left: 50, top: 50
})

canvas.add(circle, triangle);

*/ 

// testing, try to add images
var imgElement = document.getElementById('my-image');
var imgInstance = new fabric.Image(imgElement, {
  left: 100,
  top: 1,
});
canvas.add(imgInstance);
canvas.setBackgroundImage("/images/")


canvas.requestRenderAll();
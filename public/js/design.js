// Add an event listener to the "Save Changes" button
document.getElementById('saveChangesButton').addEventListener('click', function() {
    // Convert the canvas to SVG
    var svgString = canvas.toSVG();
  
    // Send the SVG data to the server
    fetch('/save-artwork', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        svgString: svgString,
      }),
    })
      .then(response => response.json())
      .then(data => {
        console.log('Server response:', data);
        // You can handle the response from the server here
      })
      .catch(error => {
        console.error('Error:', error);
      });
  });
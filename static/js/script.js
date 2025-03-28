document.querySelector('form').addEventListener('submit', function(event) {
  event.preventDefault();
  const formData = new FormData(event.target);
  fetch('/submit', {
      method: 'POST',
      body: formData
  })
  .then(response => response.text())
  .then(data => {
      alert(data);
  })
  .catch(error => {
      console.error('Error:', error);
  });
});
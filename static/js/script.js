function displayModal(content) {
    document.getElementById('resultContent').innerHTML = content;
    $('#resultModal').modal('show');
  }
  
  async function getClientIP() {
    try {
      const response = await fetch('https://api.ipify.org?format=json');
      const data = await response.json();
      return data.ip;
    } catch (error) {
      return 'Unable to retrieve IP';
    }
  }
  
  document.getElementById('phishingForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    let resultText = 'You submitted the following data:<br>';
    formData.forEach((value, key) => {
      resultText += `<strong>${key}:</strong> ${value}<br>`;
    });
  
    const ipAddress = await getClientIP();
    resultText += `<strong>IP Address:</strong> ${ipAddress}<br>`;
  
    displayModal(resultText);
  });
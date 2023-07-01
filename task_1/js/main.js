async function getResponse() {
  const response = await fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1', {
    method: 'GET',
    headers: {
      accept: 'application/json',
    },
  })
  if (!response.ok) {
    throw new Error(`Error! status: ${response.status}`)
  }
  const content = await response.json()
  console.log(content)

  // create table
  let tableContent = '<table>', ind = 0, symbol
  tableContent += '<tr><td>id</td><td>symbol</td><td>name</td></tr>'
  
  for (i in content) {
    if (content[i].symbol === 'usdt')
      symbol = '<div class=green>' + content[i].symbol + '</div>'
    else if (ind < 5)
      symbol = '<div class=blue>' + content[i].symbol + '</div>'
    else
      symbol = content[i].symbol
    ind++
    tableContent += '<tr><td>' + content[i].id + '</td><td>' + symbol + '</td><td>' + content[i].name + '</td></tr>'
  }

  tableContent += '</table>'
  document.getElementsByTagName('body')[0].innerHTML = tableContent
}

getResponse()
export function to2 (v) {
  return ('0' + v).slice(-2)
}

export function formatDate (date) {
  let d = new Date(date)
  const dayNames = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
  return `${d.getFullYear()}/${d.getMonth()+1}/${d.getDate()} (${dayNames[d.getDay()]}) `
}

export function formatTime (date) {
  let d = new Date(date)
  return `${to2(d.getHours())}:${to2(d.getMinutes())}`
}

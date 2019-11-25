
const RESIZED_WIDTH = 1024
const RESIZED_HEIGHT = 1024

const imageResizer = function (file, callback) {

  let image = new Image()

  if (typeof file === typeof {}) {
    let reader = new FileReader()
    reader.addEventListener('load', function() {
      image.src = reader.result
    }, false)
    reader.readAsDataURL(file)
  } else {
    image.src = file
  }

  image.onload = function () {
    let width, height

    let ratio = image.height / image.width
    if (image.width > image.height) {
      width = RESIZED_WIDTH
      height = RESIZED_WIDTH * ratio
    } else {
      let ratio = image.width / image.height
      width = RESIZED_HEIGHT * ratio
      height = RESIZED_HEIGHT
    }
    let canvas = document.createElement('canvas')
    canvas.setAttribute('width', width)
    canvas.setAttribute('height', height)
    let ctx = canvas.getContext('2d')
    ctx.clearRect(0, 0, width, height)
    ctx.drawImage(image, 0, 0, image.width, image.height, 0, 0, width, height)

    let base64 = canvas.toDataURL('image/jpeg')
    callback(base64)
  }
}

export default imageResizer

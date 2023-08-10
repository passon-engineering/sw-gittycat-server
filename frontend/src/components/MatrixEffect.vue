<template>
  <canvas ref="canvas" style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; z-index: -1;"></canvas>
</template>

<script>
export default {
  name: 'MatrixEffect',
  mounted() {
    this.startMatrix();
  },
  methods: {
    startMatrix() {
      let canvas = this.$refs.canvas;
      let ctx = canvas.getContext('2d');

      canvas.height = window.innerHeight;
      canvas.width = window.innerWidth;

      let matrixSymbols = 'GITTYCAT';
      matrixSymbols = matrixSymbols.split('');

      let fontSize = 20;
      let columns = canvas.width/fontSize;

      let drops = [];
      for(let i = 0; i < columns; i++)
        drops[i] = Math.floor(Math.random() * -canvas.height / fontSize);

      function draw() {
        ctx.fillStyle = 'rgba(0, 0, 0, 0.1)';  // Paint a semi-transparent rectangle over the old frame
        ctx.fillRect(0, 0, canvas.width, canvas.height);

        ctx.fillStyle = '#ff00ff';
        ctx.font = fontSize + 'px arial';

        for(let i = 0; i < drops.length; i++) {
          ctx.fillText(
            matrixSymbols[Math.floor(Math.random()*matrixSymbols.length)],
            i*fontSize,
            drops[i]*fontSize
          );

          if(drops[i]*fontSize > canvas.height && Math.random() > 0.975)
            drops[i] = 0;

          drops[i]++;
        }
      }

      setInterval(draw, 60);
    }
  }
}
</script>

<style scoped>
#matrixCanvas {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  z-index: -1;
}
</style>

<template>
  <img :src='bg' alt="" width="1920px">
  <div class="infoBorder">
    <div :class="useScroll ? 'songInfo_withAnimate':'songInfo'" ref="text">
      {{ songText }}
    </div>
  </div>
  <div class="lyricBorder">
    歌词显示在这
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
import bg from "./assets/bg.png";

const useScroll=ref(false);
const songText = ref('...');
const lyricText=ref('...');

watch(songText, (newVal)=>{
  const text=ref(null);
  if(text.value.scrollWidth>620){
    useScroll.value=true;
  }else{
    useScroll.value=false;
  }
})
</script>

<style>
.lyricBorder{
  position: fixed;
  bottom: 178px;
  height: 60px;
  width: 730px;
  left: 730px;
  line-height: 60px;
  color: white;
  font-size: 27px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
@keyframes scroll-text {
  0% {
    transform: translateX(0);
  }

  100% {
    transform: translateX(-100%);
  }
}

.songInfo_withAnimate{
  display: inline-block;
  padding-left: 100%;
  animation: scroll-text 15s linear infinite;
}

.songInfo {
  display: inline-block;
}

.infoBorder {
  position: fixed;
  bottom: 180px;
  width: 630px;
  height: 60px;
  /* background-color: red; */
  padding-left: 10px;
  /* display: flex; */
  /* align-items: center; */
  line-height: 60px;
  font-size: 40px;
  font-weight: bold;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
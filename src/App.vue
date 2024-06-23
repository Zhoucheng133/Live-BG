<template>
  <img :src='bg' alt="" width="1920px">
  <div class="infoBorder">
    <div id="text" :class="useScroll ? 'songInfo_withAnimate':'songInfo'" >
      {{ songText }}
    </div>
  </div>
  <div class="lyricBorder">
    {{lyricText}}
  </div>
</template>

<script setup>
import { nextTick, ref, watch } from "vue";
import bg from "./assets/bg.png";

const useScroll=ref(false);
const songText = ref('...');
const lyricText=ref('...');

watch(songText, async (newVal)=>{
  useScroll.value=false;
  await nextTick();
  const text=document.getElementById('text');
  if(text.scrollWidth>620){
    useScroll.value=true;
  }else{
    useScroll.value=false;
  }
})

// ws服务连接
const socket=new WebSocket('ws://localhost:9098');
socket.onmessage=function(event){
  // console.log(event.data);
  const lines=event.data.split('\n');
  songText.value=`${lines[1]} - ${lines[0]}`;
  lyricText.value=`${lines[2]}`;
}
</script>

<style>
body{
  margin: 0;
}
.lyricBorder{
  position: absolute;
  /* bottom: 178px; */
  top: 949px;
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
  position: absolute;
  /* bottom: 180px; */
  top: 950px;
  width: 630px;
  height: 60px;
  /* background-color: red; */
  padding-left: 10px;
  /* display: flex; */
  /* align-items: center; */
  line-height: 60px;
  font-size: 35px;
  font-weight: bold;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
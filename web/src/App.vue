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

<script setup lang="ts">
import { nextTick, ref, watch } from "vue";
import bg from "./assets/bg.png";
import axios from "axios";

const useScroll=ref(false);
const songText = ref('...');
const lyricText=ref('...');

watch(songText, async ()=>{
  useScroll.value=false;
  await nextTick();
  const text=document.getElementById('text');
  if(text){
    if(text.scrollWidth>620){
      useScroll.value=true;
    }else{
      useScroll.value=false;
    }
  }
  
})
// ws服务连接
const initWs=async ()=>{
  const port=(await axios.get("/api/port")).data;
  const socket=new WebSocket(`ws://localhost:${port}`);  
  socket.onopen=()=>{
    socket.send(JSON.stringify({
      'command': 'get'
    }));
  }
  socket.onmessage=function(event){
    const jsonData=JSON.parse(event.data);
    songText.value=`${jsonData.artist} - ${jsonData.title}`;
    lyricText.value=jsonData.lyric;
  }
}
initWs();
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
  left: 723px;
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
  /* padding-left: 10px; */
  left: 8px;
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
<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)

console.log("mounted!");

window.onload = function() {

}

function setProgressBar(percent: number) {
  const progressBar = document.querySelector('.progress-bar');
  console.log("setProgressBar percent: " + percent);

  if (progressBar != null) {
    const pb = progressBar as SVGTextElement
    pb.style.transform = `translateX(${percent}%)`
    // pb.style.transition = " all 1s"
  }
}

function startWS() {
  const socket = new WebSocket("ws://localhost:3333/ws/actions");

  socket.onopen = () => {
    const msg = JSON.stringify({
      type: 1,
      work_id: "e454b2cc-7620-4999-9336-cf10a41aebb6",
    });

    const h1 = document.getElementById("title");
    if (h1 != null) {
      h1.textContent = "e454b2cc-7620-4999-9336-cf10a41aebb6";
      console.log("h1: " + h1);
    }

    socket.send(msg);
  };

  socket.onmessage = (event) => {
    console.log("received from the server: ", event.data);

    const payload = JSON.parse(event.data) as ActionPayload;
    handleMessage(payload);
  };

  socket.onclose = (e: CloseEvent) => {
    console.log("socket onclose:");
    socket.onmessage = null;
  };
}

async function handleMessage(payload: ActionPayload) {
  switch (payload.type) {
    case 1:
      // const
      console.log("payload.step_number: " + payload.step_number)
      const steps = payload.step_completed ? payload.step_number + 1 : payload.step_number
      count.value = steps
      setProgressBar(100 * steps/payload.max_steps)
      break;
    default: {
      console.log("unknown payload.type: " + payload.type);
    }
  }
}

type ActionPayload = {
  type: number;
  work_id: string;
  step_number: number;
  max_steps: number;
  step_completed: boolean;
  done: boolean;
};
</script>

<template>
  <h1 id="title">{{ msg }}</h1>

  <div class="card">
    <button type="button" @click="startWS">
      <p v-if="count == 0">Start Action</p>
      <p v-else>step is {{ count }}</p>
    </button>
    <p>
      Progress Indicator
    </p>
    <div class="progress-bar__container">
      <div class="progress-bar">
      </div>
    </div>
  </div>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
      >create-vue</a
    >, the official Vue + Vite starter
  </p>
  <p>
    Install
    <a href="https://github.com/johnsoncodehk/volar" target="_blank">Volar</a>
    in your IDE for a better DX
  </p>
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}

.progress-bar__container {
  width: 80%;
  height: 2rem;
  border-radius: 2rem;
  position: relative;
  overflow: hidden;
  transition: all 0.5s;
  will-change: transform;
  box-shadow: 0 0 5px #e76f51;
}

.progress-bar {
  position: absolute;
  height: 100%;
  width: 100%;
  content: "";
  background-color: #e76f51;
  top:0;
  bottom: 0;
  left: -100%;
  border-radius: inherit;
  display: flex;
  justify-content: center;
  align-items:center;
  color: white;
  font-family: sans-serif;
}

</style>

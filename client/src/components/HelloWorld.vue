<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)

console.log("mounted!");

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
      count.value = payload.step_number
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
  step_completed: boolean;
  done: boolean;
};
</script>

<template>
  <h1 id="title">{{ msg }}</h1>

  <div class="card">
    <button type="button" @click="startWS">step is {{ count }}</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
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
</style>

import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

createApp(App).mount("#app");

// console.log("mounted!");

// export function startWS() {
//   const socket = new WebSocket("ws://localhost:3333/ws/actions");

//   socket.onopen = () => {
//     const msg = JSON.stringify({
//       type: 1,
//       work_id: "e454b2cc-7620-4999-9336-cf10a41aebb6",
//     });

//     const h1 = document.getElementById("title");
//     if (h1 != null) {
//       h1.textContent = "e454b2cc-7620-4999-9336-cf10a41aebb6";
//       console.log("h1: " + h1);
//     }

//     socket.send(msg);
//   };

//   socket.onmessage = (event) => {
//     console.log("received from the server: ", event.data);
//     const payload = JSON.parse(event.data) as ActionPayload;
//     handleMessage(payload);
//   };

//   socket.onclose = (e: CloseEvent) => {
//     console.log("socket onclose:");
//     socket.onmessage = null;
//   };
// }

// async function handleMessage(payload: ActionPayload) {
//   switch (payload.type) {
//     case 1:
//       // const
//       break;
//     default: {
//       console.log("unknown payload.type: " + payload.type);
//     }
//   }
// }

// type ActionPayload = {
//   type: number;
//   work_id: string;
//   step_number: number;
//   step_completed: boolean;
//   done: boolean;
// };

// function webSocket() {
//   // let fallback: (() => void) | undefined
//   let fallback: () => void;
//   let socket: WebSocket;

//   const socketProtocol = "ws";
//   const directSocketHost = "localhost:3333/ws/actions";

//   fallback = () => {
//     socket = setupWebSocket(socketProtocol, directSocketHost, () => {
//       console.error("failed to connect to websocket.");
//     });
//     socket.addEventListener(
//       "open",
//       () => {
//         console.info("websocket connection fallbacked");
//       },
//       { once: true }
//     );
//   };
// }

// function setupWebSocket(
//   protocol: string,
//   hostAndPath: string,
//   onCloseWithoutOpen?: () => void
// ) {
//   const socket = new WebSocket(`${protocol}://${hostAndPath}`, "vite-hmr");
//   let isOpened = false;

//   socket.addEventListener(
//     "open",
//     () => {
//       isOpened = true;
//     },
//     { once: true }
//   );

//   // Listen for messages
//   socket.addEventListener("message", async ({ data }) => {
//     handleMessage(JSON.parse(data));
//   });

//   // Ping server
//   socket.addEventListener("close", async ({ wasClean }) => {
//     if (wasClean) return;

//     if (!isOpened && onCloseWithoutOpen) {
//       onCloseWithoutOpen();
//       return;
//     }

//     console.log(`server connection lost. polling for restart...`);
//     // await waitForSuccessfulPing(protocol, hostAndPath)
//     location.reload();
//   });

//   return socket;
// }

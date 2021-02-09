<template>
  <div>
    <!-- the message's input -->
    <input id="input" type="text">
    <!-- when clicked then a websocket event will be sent to the server, at this example we registered the 'chat' -->
    <button id="sendBtn" disabled>Send</button>
    &nbsp;
    <!-- the messages will be shown here -->
    <pre>{{ outputTxt }}</pre>
  </div>
</template>

<script>
import neffos from 'neffos.js'
import { getToken } from '@/utils/auth'
export default {
  name: 'Default',
  data() {
    return {
      outputTxt: ''
    }
  },
  created() {
    this.runExample()
  },
  methods: {
    addMessage(msg) {
      this.outputTxt += (msg + '\n')
    },
    handleNamespaceConnectedConn(nsConn) {
      nsConn.emit('Hello from browser client side!')
      const token = getToken()
      nsConn.emit('Authorization', token)
      const inputTxt = document.getElementById('input')
      const sendBtn = document.getElementById('sendBtn')

      sendBtn.disabled = false
      sendBtn.onclick = () => {
        const input = inputTxt.value
        inputTxt.value = ''
        nsConn.emit('Authorization', input)
        this.addMessage('Me: ' + input)
      }
    },
    handleError(reason) {
      console.log(reason)
      window.alert('error: see the dev console')
    },
    async runExample() {
      const events = new Object()
      const that = this
      const wsURL = 'ws://192.168.50.102:5000/v1/webs/websocket'
      const token = getToken()

      events._OnNamespaceConnected = function(nsConn, msg) {
        if (nsConn.conn.wasReconnected()) {
          that.addMessage('re-connected after ' + nsConn.conn.reconnectTries.toString() + ' trie(s)')
        }

        that.addMessage('connected to namespace: ' + msg.Namespace)
        that.handleNamespaceConnectedConn(nsConn)
      }

      events._OnNamespaceDisconnect = function(nsConn, msg) {
        that.addMessage('disconnected from namespace: ' + msg.Namespace)
      }

      events.chat = function(nsConn, msg) { // "chat" event.
        that.addMessage(msg.Body)
      }
      try {
        // You can omit the "default" namespace and simply define only Events,
        // the namespace will be an empty string"",
        // however if you decide to make any changes on
        // this example make sure the changes are reflecting inside the ../server.go file as well.
        //
        // At "wsURL" you can put the relative URL if the client and server
        // hosted in the same address, e.g. "/echo".
        const conn = await neffos.dial(wsURL, { default: events }, {
          // if > 0 then on network failures it tries to reconnect every 5 seconds, defaults to 0 (disabled).
          reconnect: 5000,
          // custom headers:
          headers: {
            'Authorization': getToken()
          }
        })
        conn.connect('default')
      } catch (err) {
        that.handleError(err)
      }
    }

  }
}
</script>

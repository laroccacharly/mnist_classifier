export default function (state, action) {
    console.log("Received message type", action.type)

    switch (action.type) {
        case "GET_SAMPLES": {
            // add id to samples
            const samples = action.samples.map((s, i) => ({...s, id: i}))
            return {...state, samples: samples}
        }
        default : {
            return state
        }
    }
}
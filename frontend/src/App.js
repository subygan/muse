import logo from './logo.svg';
import './App.css';
import { useEffect,useState } from 'react'

function App() {
  const [url_gen, setUrl_gen] = useState('')
  console.log(url_gen )

  const axios = require('axios')

  const get_link = async () => {
    try {
      axios.get('https://83c5-49-207-209-240.ngrok.io/')
      .then(res=>setUrl_gen(res.data.Url))
    } catch (error) {
      console.error(error)
    }
  }


  useEffect(() => {
    get_link()

  }, [])

  useEffect(() => {

  }, [url_gen])

  return (
    <div className="App">
     

      {url_gen ? 
       <audio controls>
        <source src={url_gen} type="audio/mpeg" />
        Your browser does not support the audio element.
      </audio> : null}
    </div>
  );
}

export default App;

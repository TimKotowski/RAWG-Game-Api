import React, {useEffect, useState} from "react";
import axios from "axios"
import Loading from './loading.gif'
import './App.css';

function App() {
  const [data, setData] = useState([]);
  const [isLoading, setLoading] = useState(false);


  useEffect(() => {
    const fetchData = async () => {
      setLoading(true)
      const result = await axios.get('/api/v1/users');
      setData(result.data);
      setLoading(false)
    };
    fetchData();
  }, []);


  const {results} = data
  console.log("sds", results)
  return (
    <div className="App">
        {!isLoading && results && results.length > 0  ? (
          <div>
             {results.map(item => (
              <div key ={item.id}>
              <h5>{item.name}</h5>
              <img  className="photo" alt="sd" src={item.background_image} />
              </div>
          ))}
          </div>
      ) : (
        <div id="main" >
        <img className="p" alt="giff" src={Loading} />
        </div>
      )}
      </div>
      );
    }

export default App;

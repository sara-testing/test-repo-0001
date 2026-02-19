import logo from './logo.svg';
import './App.css';
import { useState } from "react";

function App() {
  const [personalInformation, setPersonalInformation] = useState({
    name: "",
    surname:"",
    id: "",
    address: ""
  })
  
  return (
    <div className="App">
      <header className="App-header">Your personal information</header>
        <form>
          <input type="text" label="name" placeholder="name" onChange={e => setPersonalInformation({...personalInformation, name: e.target.value})}/>
          <input type="text" label="surname" placeholder="surname" onChange={e => setPersonalInformation({...personalInformation, surname: e.target.value})}/>
          <input type="text" label="id" placeholder="your personal id" onChange={e => setPersonalInformation({...personalInformation, id: e.target.value})}/>
          <input type="text" label="address" placeholder="your home address" onChange={e => setPersonalInformation({...personalInformation, address: e.target.value})}/>
         <button type="submit">save your information</button> 
        </form>    
    </div>
  );
}

export default App;

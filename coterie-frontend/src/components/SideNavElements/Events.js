import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Event from './Event';
import NewEvent from './NewEvent';


const Events = (props) => {
  const [eventsList, setEventsList] = useState(null);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const addEvent = (eventObj) => {
    axios.post(url, eventObj)
    .then((response) => {
      setErrorMessage(`Event ${eventObj["name"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add mem: ${errorMessage}`);
    })
  }
  
  
    //toggle visibility of new member component
    const toggleFormVisibility = () => {
      setVisibility(!visibility);
      return;
    }

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/events`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setEventsList(list);
      })
      .catch((error) => {
        console.log(`There was an error retrieving events: ${error}`)
      });
  },[url])

  let eventComponents = undefined
  if (eventsList !== null && eventsList.length > 0) {
    eventComponents = eventsList.map((e) => {
    return(
      <Event
      key={e.id}
      id={e.id}
      name={e.name}
      occasion={e.occasion}
      date={e.date}
      description={e.description}
      orgID={e.organization_id}
      />
    )
  })
}


  return (
    <div className="container">
      <div className="d-flex list-group">
      <div className="d-flex py-2 justify-content-between">
        <p className={ eventComponents !== undefined ? "hidden" : "open-sans" }>You haven't created any events yet.</p>
          <button className="btn btn-secondary" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        </div>
        <NewEvent 
        orgID={props.orgID}
        visibility={visibility}
        addEventCallback={addEvent}
        onSubmitCallback={toggleFormVisibility}
        />
        {eventComponents}
       
      </div>
    </div>
  )
}

export default Events;


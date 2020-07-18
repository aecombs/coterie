import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Event from './Event';


const Events = (props) => {
  const [eventsList, setEventsList] = useState(null);

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
      createdAt={e.created_at}
      orgID={e.organization_id}
      />
    )
  })
} else {
  eventComponents = <p className="open-sans">It looks like you don't have any events, yet!</p>
}


  return (
    <div className="container">
      <h4 className="text-left w-100">All Events</h4>
      <div className="d-flex list-group">
        {eventComponents}
      </div>
    </div>
  )
}

export default Events;


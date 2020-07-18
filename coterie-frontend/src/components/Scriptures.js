import React, { useState, useEffect } from 'react';
import Scripture from './Scripture';
import axios from 'axios';

const Scriptures = (props) => {
  const [scripturesList, setScripturesList] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/scriptures`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setScripturesList(list);
      })
      .catch((error) => {
        console.log(`There was an retrieving scriptures: ${error}`)
      });
  },[url])
  
  let scriptureComponents = undefined
  if (scripturesList !== null && scripturesList.length > 0) {
    scriptureComponents = scripturesList.map((scrip) => {
    return (<Scripture 
          key={scrip.id}
          id={scrip.id}
          name={scrip.name}
          orgID={scrip.organization_id}
        />)
    })
  } else {
    scriptureComponents = <p className="open-sans">It looks like you don't have any scriptures, yet!</p>
  };

  return (
    <div className="">
      <p className="open-sans">Scriptures</p>
      {scriptureComponents}
    </div>
  )
}

export default Scriptures;
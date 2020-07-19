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
        console.log(`There was an error retrieving scriptures: ${error}`)
      });
  },[url])
  
  let scriptureComponents = undefined
  if (scripturesList !== null && scripturesList.length > 0) {
    scriptureComponents = scripturesList.map((scrip) => {
    return (
      <Scripture 
      key={scrip.id}
      id={scrip.id}
      name={scrip.name}
      orgID={scrip.organization_id}
      />)
    })
  } else {
    scriptureComponents = <p className="open-sans">Add your important scriptures here!</p>
  };

  return (
    <div className="container">
      <h4 className="text-left w-100">All Scriptures</h4>
      <div className="d-flex list-group">
        {scriptureComponents}
      </div>
    </div>
  )
}

export default Scriptures;
import React, { useState, useEffect } from 'react';
import Scripture from './Scripture';
import ScriptureForm from './ScriptureForm';
import axios from 'axios';

const Scriptures = (props) => {
  const [scripturesList, setScripturesList] = useState(null);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/scriptures`

  const updateURL = `http://localhost:3000/scriptures`

  const addScripture = (scripObj) => {
    //remove unnecessary id property
    delete scripObj["id"];

    axios.post(url, scripObj)
    .then((response) => {
      setErrorMessage(`Scripture ${scripObj["name"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add scripture: ${errorMessage}`);
    })
  }

  const updateScripture = (scripObj) => {
    axios.put(`${updateURL}/${scripObj.id}`, scripObj)
    .then((response) => {
      setErrorMessage(`Scripture ${scripObj["name"]} was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add scripture: ${errorMessage}`);
    })
  }

   //toggle visibility of scripture form component
   const toggleFormVisibility = () => {
    setVisibility(!visibility);
    return;
  }


  //retrieve scriptures
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
      updateScriptureCallback={updateScripture}
      />)
    })
  }

  return (
    <div className="container">
      <div className="d-flex list-group">
        <div className="d-flex py-2 justify-content-between">
          <h4>Scriptures</h4>
          <button className="btn btn-secondary" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        </div>
        <p className={ scriptureComponents !== undefined ? "hidden" : "open-sans" }>You haven't created any scriptures yet.</p>
        <ScriptureForm 
        orgID={props.orgID}
        visibility={visibility}
        addScriptureCallback={addScripture}
        onSubmitCallback={toggleFormVisibility}
        />
        {scriptureComponents}
      </div>
    </div>
  )
}

export default Scriptures;
import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import ScriptureForm from './ScriptureForm';
import Chapter from './Chapter';
import ChapterForm from './ChapterForm';
import axios from 'axios';

const Scripture = (props) => {
  const [chaptersList, setChaptersList] = useState(null);
  const [updateChapterMode, setUpdateChapterMode] = useState(false);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);
  const [updateScriptureMode, setUpdateScriptureMode] = useState(false);

  const updateScripture = (scripObj) => {
    props.updateScriptureCallback(scripObj)
  }

  const setUpdateScripture = () => {
    setUpdateScriptureMode(!updateScriptureMode);
  }

  const setUpdateChapter = () => {
    setUpdateChapterMode(!updateChapterMode);
  }


  const url = `http://localhost:3000/scriptures/${props.id}/chapters`

  const updateChapterURL = `http://localhost:3000/chapters`

  const addChapter = (chapObj) => {
    //remove unnecessary id property
    delete chapObj["id"];

    axios.post(url, chapObj)
    .then((response) => {
      setErrorMessage(`Chapter ${chapObj["name"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add scripture: ${errorMessage}`);
    })
  }

  const updateChapter = (chapObj) => {
    axios.put(`${updateChapterURL}/${chapObj.id}`, chapObj)
    .then((response) => {
      setErrorMessage(`Chapter ${chapObj["name"]} was updated`);
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


  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setChaptersList(list);
      })
      .catch((error) => {
        console.log(`There was an error retrieving chapters: ${error}`)
      });
  },[url])
  
  let chapterComponents = undefined
  if (chaptersList !== null && chaptersList.length > 0) {
    chapterComponents = chaptersList.map((chap) => {
    return (
      <Chapter 
      key={chap.id}
      id={chap.id}
      name={chap.name}
      text={chap.text}
      position={chap.position}
      scripID={chap.scripture_id}
      updateChapterCallback={updateChapter}
      />)
    })
  }


  return (
    <section>

      <button onClick={setUpdateChapter} className={ updateChapterMode ? "hidden" : "btn list-group-item list-group-item-action"}>
        <div className="card-body row justify-content-between">
          <p className="card-title font-weight-bolder text-left">{props.name}</p>
          <button className="btn btn-secondary" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        </div>

        <p className={ chapterComponents !== undefined ? "hidden" : "open-sans" }>There are no chapters in this scripture...</p>
        <ChapterForm 
        orgID={props.orgID}
        visibility={visibility}
        scripLength={chapterComponents.length}
        addChapterCallback={addChapter}
        onSubmitCallback={toggleFormVisibility}
        />
        {chapterComponents}

      </button>
      <ScriptureForm 
        id={props.id}
        name={props.name}
        orgID={props.orgID}
        visibility={updateScriptureMode}
        addScriptureCallback={updateScripture}
        updateScriptureCallback={updateScripture}
        onSubmitCallback={setUpdateScripture}
        />
    </section>
  )
}

Scripture.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  orgID: PropTypes.number
};

export default Scripture;
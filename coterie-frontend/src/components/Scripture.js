import React, { useState, useEffect } from 'react';
import Chapter from './Chapter';
import axios from 'axios';

const Scripture = (props) => {
  const [chaptersList, setChaptersList] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/scriptures/${props.id}/chapters`

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
      />)
    })
  } else {
    chapterComponents = <p className="open-sans">This scripture still needs some chapters.</p>
  };


  return (
    <div className="">
      <h6 className="text-left w-100">{props.name}</h6>
      <div className="">
        {chapterComponents}
      </div>
    </div>
  )
}

export default Scripture;
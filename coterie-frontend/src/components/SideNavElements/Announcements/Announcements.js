import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Announcement from './Announcement';
import AnnouncementForm from './AnnouncementForm';


const Announcements = (props) => {
  const [announcementsList, setAnnouncementsList] = useState(null);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}/organizations/${props.orgID}/announcements`

  const updateURL = `${process.env.REACT_APP_API_BASE_URL}/announcements`

  const addAnnouncement = (annObj) => {
    //remove unnecessary id property
    delete annObj["id"];

    axios.post(url, annObj)
    .then((response) => {
      setErrorMessage(`Announcement ${annObj["header"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add announcement: ${errorMessage}`);
    })
  }

  const updateAnnouncement = (annObj) => {
    axios.put(`${updateURL}/${annObj.id}`, annObj)
    .then((response) => {
      setErrorMessage(`Announcement ${annObj["header"]} was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to update announcement: ${errorMessage}`);
    })
  }

   //toggle visibility of announcement form component
   const toggleFormVisibility = () => {
    setVisibility(!visibility);
    return;
  }

  //Get all announcements
  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setAnnouncementsList(list);
      })
      .catch((error) => {
        setErrorMessage(error);
        console.log(`There was an error retrieving announcements: ${error}`)
      });
  },[url])

  let announcementComponents = undefined
  if (announcementsList !== null && announcementsList.length > 0) {
    announcementComponents = announcementsList.map((ann) => {
    return(
      <Announcement
      key={ann.id}
      id={ann.id}
      header={ann.header}
      description={ann.description}
      date={ann.date}
      orgID={ann.organization_id}
      updateAnnouncementCallback={updateAnnouncement}
      />
    )
  })
}


  return (
    <div className="container">
      <div className="d-flex list-group">
        <div className="d-flex py-2 justify-content-between">
          <h4>Announcements</h4>
          <button className="btn btn-secondary" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        </div>
        <p className={ announcementComponents !== undefined ? "hidden" : "open-sans" }>You haven't created any announcements yet.</p>
        <AnnouncementForm 
        orgID={props.orgID}
        visibility={visibility}
        addAnnouncementCallback={addAnnouncement}
        onSubmitCallback={toggleFormVisibility}
        />
        {announcementComponents}
      </div>
    </div>
  )
}

export default Announcements;
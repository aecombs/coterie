import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Announcement from './Announcement';


const Announcements = (props) => {
  const [announcementsList, setAnnouncementsList] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/announcements`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setAnnouncementsList(list);
      })
      .catch((error) => {
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
      text={ann.text}
      date={ann.date}
      orgID={ann.organization_id}
      />
    )
  })
} else {
  announcementComponents = <p className="open-sans">You haven't made any announcements yet.</p>
}


  return (
    <div className="container">
      <h4 className="text-left w-100">All Announcements</h4>
      <div className="d-flex list-group">
        {announcementComponents}
      </div>
    </div>
  )
}



export default Announcements;
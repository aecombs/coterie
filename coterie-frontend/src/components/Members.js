import React, {useState, useEffect} from 'react';
import Member from './Member';
import axios from 'axios';

const Members = (props) => {
  const [membersList, setMembersList] = useState(null);

  const url = "http://localhost.com/3000/organization/members"

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.headers;
        setMembersList(list);
      })
      .catch((error) => {
        console.log(`There was an retrieving members: ${error}`)
      });
  },[url])

  let memberComponents = undefined
  if (membersList !== null && membersList.length > 0) {
    memberComponents = membersList.map((mem) => {
    return(
      <Member
      key={mem.id}
      id={mem.id}
      name={mem.name}
      birthdate={mem.birthdate}
      class={mem.class}
      email={mem.email}
      fundsRaised={mem.funds_raised}
      orgID={mem.organization_id}
      />
    )
  })
} else {
  memberComponents = <p className="open-sans">It looks like you don't have any members, yet!</p>
}


  return (
    <div>
      <p className="open-sans">Members</p>
      <div className="">{memberComponents}</div>
    </div>
  )
}

export default Members;


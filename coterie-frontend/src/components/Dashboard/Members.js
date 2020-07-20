import React, {useState, useEffect} from 'react';
import Member from './Member';
import MemberForm from './MemberForm';
import axios from 'axios';

const Members = (props) => {
  const [membersList, setMembersList] = useState(null);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);
  
  //get all members
  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}/organizations/${props.orgID}/members`;

  const addMember = (memObj) => {
    delete memObj["id"];
    
    axios.post(url, memObj)
    .then((response) => {
      setErrorMessage(`Member ${memObj["name"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add mem: ${errorMessage}`);
    })
  }

  const updateURL = `${process.env.REACT_APP_API_BASE_URL}/members`

  const updateMember = (memObj) => {
    axios.put(`${updateURL}/${memObj.id}`, memObj)
    .then((response) => {
      setErrorMessage(`Member ${memObj["header"]} was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to update member: ${errorMessage}`);
    })
  }


  //toggle visibility of new member component
  const toggleFormVisibility = () => {
    setVisibility(!visibility);
    return;
  }


  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setMembersList(list);
      })
      .catch((error) => {
        setErrorMessage(error);
        console.log(`There was an error retrieving members: ${error}`)
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
      submitMemberCallback={updateMember}
      />
    )
  })
} else {
  memberComponents = <p className="open-sans">It looks like you don't have any members, yet!</p>
}


  return (
    <section className="w-100">
      <div className="">
        <button className="btn btn-secondary float-right mb-2" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        <MemberForm 
        orgID={props.orgID}
        visibility={visibility}
        submitMemberCallback={addMember}
        onSubmitCallback={toggleFormVisibility}
        />
      </div>
      <h6 className="mt-3 w-100">Followers</h6>
      {/* <table className="table table-hover table-light">
        <thead className="thead-light text-left">
          <tr>
            <th>Name</th>
            <th>Class</th>
            <th>Funds Raised</th>
            <th>Email</th>
            <th>Birthdate</th>
          </tr>
        </thead>
        <tbody> */}
          {memberComponents}
        {/* </tbody>
      </table> */}
    </section>
  )
}

export default Members;


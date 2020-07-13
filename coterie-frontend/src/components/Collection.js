import React from 'react';
import Announcement from './Announcement';
import Event from './Event';
import Holiday from './Holiday';
import Scripture from './Scripture';
import Member from './Member';

const Collection = (props) => {
  return (
    <div>
      <p className="baloo-da">{props.selection}</p>
    </div>
  )
}

export default Collection;
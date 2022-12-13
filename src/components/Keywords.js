import { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";
import { BrowserRouter as Router, Route, Link, redirect } from "react-router-dom";
import { Button } from "@chakra-ui/react";
import { format, formatDistance, formatRelative, subDays } from 'date-fns'
import { json } from "react-router-dom";
import Keyword from "./Keyword";
// import axios from 'axios'

const Keywords = (props) => {
  // const navigate = useNavigate();


  return (
    <div>
      <h2>📘  Keywords History</h2>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Keyword</th>
            <th>Description</th>
            <th>......</th>
          </tr>
        </thead>
        <tbody>
          {props.keywords.map((k, index) => (
            <tr key={index}>
              <td><Link to={`/keywords/${k.ID}`}>{k.Word}</Link></td>
              {/* <td><Route path={`/keywords/${k.ID}`}><Keyword /><Route/></td> */}


              <td>{k.Description}</td>
              <td>
                {k.ImageUrl !== "" &&
                  <div className="mb-3">
                    {/* <img src={`${k.ImageUrl}`} alt="generated-img" /> */}
                    <img src={`https://res.cloudinary.com/dokzsbu2v/image/upload/v1670479204/development/images_xa8j85.png`} alt="generated-img" />
                  </div>
                }

              </td>
              {/* <td>  {format(Date.parse(k.CreatedAt),"yyyy/M/d HH:mm")}  </td> */}
              <td><Button onClick={() => props.confirmDelete(k.ID)} colorScheme={"gray"}  variant={"outline"}>
                🗑️ Delete
                </Button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default Keywords;

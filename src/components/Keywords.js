// import { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";
// import { Link, redirect } from "react-router-dom";
import { Button } from "@chakra-ui/react";
// import axios from 'axios'

const Keywords = (props) => {
  // const navigate = useNavigate();
  // const [keywords, setKeywords] = useState([]);



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
              <td>{k.Word}
                {/* <Link to={`/keywords/${k.ID}`}>
                                    {k.word}
                                </Link> */}
              </td>
              <td>{k.Description}</td>
              <td>
                {k.ImageUrl !== "" &&
                  <div className="mb-3">
                    {/* <img src={`${k.ImageUrl}`} alt="generated-img" /> */}
                    <img src={`https://res.cloudinary.com/dokzsbu2v/image/upload/v1670479204/development/images_xa8j85.png`} alt="generated-img" />
                  </div>
                }

              </td>
              <td>{k.CreatedAt}</td>
              <td><Button onClick={() => props.confirmDelete(k.ID)} colorScheme={"gray"}  variant={"outline"}>
            🗑️ Delete
          </Button></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default Keywords;

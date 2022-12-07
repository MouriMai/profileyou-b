import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Keywords = () => {
  const [keywords, setKeywords] = useState([]);

  useEffect(() => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    }

    fetch(`http://localhost:8080/keywords`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setKeywords(data);
      })
      .catch(err => {
        console.log(err);
      })

  }, []);

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
          {keywords.map((k, index) => (
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
                    <img src={`https://image.tmdb.org/t/p/w200/${k.ImageUrl}`} alt="generated-img" />
                  </div>
                }

              </td>
              <td>{k.CreatedAt}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default Keywords;

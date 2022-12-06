import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Keywords = () => {
    const [keywords, setKeywords] = useState([]);

    useEffect( () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`http://localhost:8080/keywords`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                console.log(data.lists);
                setKeywords(data.lists);
            })
            .catch(err => {
                console.log(err);
            })

    }, []);

    return(
        <div>
            <h2>Keywords</h2>
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
                    {keywords.map((k) => (
                        <tr key={k.Word}>
                            <td>{k.Word}
                                <Link to={`/keywords/${k.ID}`}>
                                    {k.word}
                                </Link>
                            </td>
                            <td>{k.Description}</td>
                            <td>{k.Image_url}</td>
                        </tr>    
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default Keywords;
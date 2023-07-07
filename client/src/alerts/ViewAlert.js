import React, { useState, useEffect } from 'react'
import { API_URL } from '../config';
import { useParams, Link } from "react-router-dom";


export default function ViewAlert() {
  let { id } = useParams();
  const [alert, setAlert] = useState({})
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    const fetchAlert = async () => {
      setLoading(true)
      try {
        const response = await fetch(`${API_URL}/alerts/${id}`);
        const json = await response.json();
        setAlert(json.data);
        setLoading(false)
      } catch (error) {
        console.log("error", error);
        setLoading(false)
      }
    };

    fetchAlert();
  }, [id]);

  return (
    <div>
      {!loading ?
        <div className="flex justify-center">
          <div className="lg:w-1/3 w-full">
            <div className="p-10">
              <div className="mb-10 flex items-center justify-between">
                <Link to="/"><h1 className="font-bold">Go back</h1></Link>
              </div>
              <div className="bg-slate-100 rounded-lg px-5">
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Type</div>
                  <div className="text-slate-800 font-medium">{alert.type}</div>
                </div>
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Description</div>
                  <div className="text-slate-800 font-medium">{alert.description}</div>
                </div>
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Created At</div>
                  <div className="text-slate-800 font-medium">{alert.created}</div>
                </div>
                <div className="flex py-4">
                  <div className="mr-4 text-slate-400">Country</div>
                  <div className="text-slate-800 font-medium">{alert.country}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        : ''}
    </div>
  )
}

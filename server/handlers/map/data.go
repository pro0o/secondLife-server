package osm

import "math"

const JsonData = `{
    "orphanages": [
        {
            "name": "Child Nepal",
            "longitude": 27.7182, 
            "latitude": 85.3513,   
            "location": "Chuchhepati, Kathmandu"
        },
        {
            "name": "Mission Himalaya Eco Home Orphanage",
            "longitude": 27.576042,
            "latitude": 85.553798,
            "location": "Shankhu Patichaur, Dhulikhel"
        },
        {
            "name": "The Orphans Homes",
            "longitude": 27.653026,
            "latitude": 85.313180,
            "location": "Lalitpur"
        },
        {
            "name": "Nepal Children's Organization",
            "longitude": 27.713866, 
            "latitude": 85.330514,
            "location": "P87J+H66, Kathmandu 44600"
        },
        {
            "name": "The Himalayan Innovative Society",
            "longitude": 27.730764,
            "latitude":  85.335081,
            "location": "P8JP+82J, Lamingtan Marg, Kathmandu 44600"
        }
    ],
    "food": [],
    "recycleCentre": [],
    "ngos": [
    {
        "name": "Bharosa Organization (NGO)",
        "longitude": 27.685743, 
        "latitude":  85.317661,   
        "location": "M8P9+93C, Lalitpur 44600"
    },
    {
        "name": "Safetyknot Nepal",
        "longitude": 27.678218, 
        "latitude": 85.315429,
        "location": "M8H8+85G, Aaksheswor Mahavihar Galli Uttar, Lalitpur 44600"
    },
    {
        "name": "NGO Federation of Nepal",
        "longitude": 27.684569, 
        "latitude":  85.328950,
        "location": "M8MH+VHJ, बुद्ध मार्ग, Kathmandu 44614"
    },
    {
        "name": "Fundraising For NGOs",
        "longitude": 27.674923, 
        "latitude":  85.341831,
        "location": "M8FR+VPX, Kriti Marg, Kathmandu 44600"
    }
        
    ],
    "clothes": [{
        "name": "Affordable Thriftstore Nepal",
        "longitude": 27.692951, 
        "latitude":  85.322037,   
        "location": "Maitighar Mandala, Kathmandu 44600"
    },
    {
        "name": "Deal and steal thrift store",
        "longitude": 27.576042,
        "latitude": 85.553798,
        "location": "Kathmandu 44600"
    },
    {
        "name": "Sukhawati Charity Store",
        "longitude": 27.677978,
        "latitude":  85.307789,
        "location": "M8H5+H44, Kathmandu 44700"
    },
    {
        "name": "Thrift Finds Nepal",
        "longitude": 27.657836,
        "latitude":  85.324800,
        "location": "P87J+H66, Kathmandu 44600"
    }
    ],
    "E-Waste": [],
    "Plastic": [
        {
            "name": "Mahakali Plastic recycler",
            "longitude": 27.661109,  
            "latitude":  85.365868,  
            "location": "M968+G7F, Anantalingeshwar 44600"
        },
        {
            "name": "Deal and steal thrift store",
            "longitude": 27.576042,
            "latitude": 85.553798,
            "location": "Kathmandu 44600"
        },
        {
            "name": "Sukhawati Charity Store",
            "longitude": 27.677978,
            "latitude":  85.307789,
            "location": "M8H5+H44, Kathmandu 44700"
        },
        {
            "name": "Thrift Finds Nepal",
            "longitude": 27.657836,
            "latitude":  85.324800,
            "location": "P87J+H66, Kathmandu 44600"
        }
    ],
    "Glass": [],
    "Paper": [],
    "all":[
        {
            "name": "Doko Recyclers",
            "longitude": 27.687096, 
            "latitude":  85.371652,   
            "location": "M9PC+WP Bhaktapur"
        },
        {
            "name": "Phohor Mohor",
            "longitude": 27.669465, 
            "latitude": 85.363887,
            "location": "Bhaktapur 44600"
        },
        {
            "name": "3R solution pvt.ltd",
            "longitude": 27.698607,  
            "latitude": 85.386398,
            "location": "M9XP+FH Kathmandu"
        }
    ]
}`

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371

	lat1Rad := degToRad(lat1)
	lon1Rad := degToRad(lon1)
	lat2Rad := degToRad(lat2)
	lon2Rad := degToRad(lon2)

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

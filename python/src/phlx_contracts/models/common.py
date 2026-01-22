"""Common geographic types used across Phylax models"""

from typing import List, Literal
from pydantic import BaseModel, Field


class GeoPoint(BaseModel):
    """
    GeoPoint represents a geographical point in GeoJSON format.
    Used for MongoDB geospatial queries.
    """

    type: Literal["Point"] = "Point"
    coordinates: List[float] = Field(
        ..., description="[longitude, latitude]", min_length=2, max_length=2
    )


class GeoLocation(BaseModel):
    """
    GeoLocation represents geographic coordinates in GeoJSON format.
    Used for mission and team locations.
    """

    type: Literal["Point"] = "Point"
    coordinates: List[float] = Field(
        ..., description="[longitude, latitude]", min_length=2, max_length=2
    )


class Coordinate(BaseModel):
    """
    Coordinate represents a GPS coordinate.
    Used for location boundaries and areas.
    """

    latitude: float
    longitude: float

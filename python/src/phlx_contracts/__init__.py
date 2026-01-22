"""
Phylax Contracts - Shared type contracts for Phylax microservices

This package provides type-safe contracts for data models and events
across all Phylax microservices.
"""

__version__ = "1.0.0"

from . import models
from . import events

__all__ = ["models", "events"]
